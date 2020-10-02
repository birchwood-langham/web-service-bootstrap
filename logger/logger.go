package logger

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/birchwood-langham/web-service-bootstrap/config"
)

func ZapConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

var once sync.Once
var syncer zapcore.WriteSyncer
var core zapcore.Core
var log *zap.Logger

func ZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(ZapConfig())
}

func ZapWriter(writer io.Writer) zapcore.WriteSyncer {
	once.Do(func() {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout))
	})

	return syncer
}

func LumberjackLogger(fileName string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func DefaultLumberjackLogger() *lumberjack.Logger {
	return LumberjackLogger(
		viper.GetString(config.LogFilePathKey),
		viper.GetInt(config.LogFileMaxSize),
		viper.GetInt(config.LogFileMaxBackups),
		viper.GetInt(config.LogFileMaxAge),
		viper.GetBool(config.LogFileCompress),
	)
}

func New(level zapcore.Level, writer io.Writer) *zap.Logger {
	once.Do(func() {
		core = zapcore.NewCore(ZapEncoder(), ZapWriter(writer), level)
		log = zap.New(core, zap.AddCaller())
	})

	defer func() {
		_ = log.Sync()
	}()

	return log
}

func ApplicationLogLevel() zapcore.Level {
	var level zapcore.Level

	switch strings.ToUpper(viper.GetString(config.LogLevelKey)) {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "FATAL":
		level = zapcore.FatalLevel
	case "PANIC":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	return level
}

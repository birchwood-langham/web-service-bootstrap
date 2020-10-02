package logger

import (
	"errors"
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

var once sync.Once
var syncer zapcore.WriteSyncer
var core zapcore.Core
var log *zap.Logger

// ZapConfig returns the bootstrap default zap configuration
func ZapConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

// ZapEncoder creates a console encoder providing the default zap configuration
func ZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(ZapConfig())
}

// ZapWriter initializes and returns a multi-sync writer
// with the given writer and stdout
// If the writer has already been initialized, it will return
// the initialized writer
func ZapWriter(writer io.Writer) zapcore.WriteSyncer {
	if syncer == nil {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout))
	}

	return syncer
}

// LumberjackLogger creates a lumberjack logger with the given parameters
func LumberjackLogger(fileName string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

// ZapCore returns the zap core used by the logger
func ZapCore() (zapcore.Core, error) {
	if core == nil {
		return nil, errors.New("ZapCore has not been initialized")
	}

	return core, nil
}

// DefaultLumberjackLogger returns the lumberjack logger using default settings
// provided in the application settings
func DefaultLumberjackLogger() *lumberjack.Logger {
	return LumberjackLogger(
		viper.GetString(config.LogFilePathKey),
		viper.GetInt(config.LogFileMaxSize),
		viper.GetInt(config.LogFileMaxBackups),
		viper.GetInt(config.LogFileMaxAge),
		viper.GetBool(config.LogFileCompress),
	)
}

// New initializes the zap core and logger for the application
// If the core and logger has already been initialized, New returns
// the existing logger
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

// ApplicationLogLevel returns the log level defined in the
// application configuration file
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

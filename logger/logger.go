package logger

import (
	"os"
	"strings"

	"github.com/birchwood-langham/web-service-bootstrap/config"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

func ZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(ZapConfig())
}

func ZapWriter() zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:   viper.GetString(config.LogFilePathKey),
		MaxSize:    viper.GetInt(config.LogFileMaxSize),
		MaxBackups: viper.GetInt(config.LogFileMaxBackups),
		MaxAge:     viper.GetInt(config.LogFileMaxAge),
		Compress:   viper.GetBool(config.LogFileCompress),
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(logger), zapcore.AddSync(os.Stderr))
}

func New(level zapcore.Level) *zap.Logger {
	core := zapcore.NewCore(ZapEncoder(), ZapWriter(), level)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	return logger
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

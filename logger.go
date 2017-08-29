package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Log ...
var (
	logStdOut *zap.Logger
	logStdErr *zap.Logger
	isVerbose bool
)

func init() {
	InitLogger(false)
}

// InitLogger : (optional)
func InitLogger(verbose bool) {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.EncoderConfig.MessageKey = "message"
	conf.EncoderConfig.TimeKey = "timestamp"

	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)
	conf.Level = level

	conf.OutputPaths = []string{"stdout"}
	logStdOut, _ = conf.Build()

	conf.OutputPaths = []string{"stderr"}
	logStdErr, _ = conf.Build()

	isVerbose = verbose
}

//Error logs errors into std err
func Error(msg string, fields ...zapcore.Field) {
	logStdErr.Error(msg, fields...)
}

//Debug logs info into std out
func Debug(msg string, fields ...zapcore.Field) {
	if isVerbose {
		logStdOut.Debug(msg, fields...)
	}
}

//Info logs info into std out
func Info(msg string, fields ...zapcore.Field) {
	logStdOut.Info(msg, fields...)
}

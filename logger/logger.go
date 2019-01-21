package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger
var logger *zap.Logger
var logLevel = zap.NewAtomicLevel()

// Level logging priority, higher is more important
type Level zapcore.Level

const (
	// DebugLevel debug message, usually disabled in production
	DebugLevel Level = Level(zapcore.DebugLevel)
	// InfoLevel info message
	InfoLevel Level = Level(zapcore.InfoLevel)
	// WarnLevel warning message
	WarnLevel Level = Level(zapcore.WarnLevel)
	// ErrorLevel error message
	ErrorLevel Level = Level(zapcore.ErrorLevel)
	// FatalLevel fatal message, log, then panic
	FatalLevel Level = Level(zapcore.FatalLevel)
	// PanicLevel panic message, log, then call os.Exit(0)
	PanicLevel Level = Level(zapcore.PanicLevel)
)

func init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		os.Stderr,
		logLevel,
	)
	logger = zap.New(core)
	sugar = logger.Sugar()
}

// SetLevel alters the logging level
func SetLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

// Debugw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues)
}

// Panicw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	sugar.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message.
func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message.
func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sugar.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sugar.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugar.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sugar.Error(args)
}

// Panic uses fmt.Sprint to construct and log a message.
func Panic(args ...interface{}) {
	sugar.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message.
func Fatal(args ...interface{}) {
	sugar.Fatal(args)
}

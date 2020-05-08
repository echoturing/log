package log

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.SugaredLogger
	once          sync.Once
)

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.OutputPaths = []string{"stdout"}
	config.Sampling = nil
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	_logger, _ := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
	defaultLogger = _logger.Sugar()
}

func UseProductConfig() {
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout"}
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		_logger, _ := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
		defaultLogger = _logger.Sugar()
	})
}

func Sync() {
	_ = defaultLogger.Sync()
}

func DebugWithContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	requestIDUserID := FromContext(ctx)
	defaultLogger.Debugw(msg, append([]interface{}{"x-request-id", requestIDUserID.RequestID, "x-user-id", requestIDUserID.UserID}, keysAndValues...)...)
}

func InfoWithContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	requestIDUserID := FromContext(ctx)
	defaultLogger.Infow(msg, append([]interface{}{"x-request-id", requestIDUserID.RequestID, "x-user-id", requestIDUserID.UserID}, keysAndValues...)...)
}

func WarnWithContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	requestIDUserID := FromContext(ctx)
	defaultLogger.Warnw(msg, append([]interface{}{"x-request-id", requestIDUserID.RequestID, "x-user-id", requestIDUserID.UserID}, keysAndValues...)...)
}

func ErrorWithContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	requestIDUserID := FromContext(ctx)
	defaultLogger.Errorw(msg, append([]interface{}{"x-request-id", requestIDUserID.RequestID, "x-user-id", requestIDUserID.UserID}, keysAndValues...)...)
}

func PanicWithContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	requestIDUserID := FromContext(ctx)
	defaultLogger.Panicw(msg, append([]interface{}{"x-request-id", requestIDUserID.RequestID, "x-user-id", requestIDUserID.UserID}, keysAndValues...)...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	defaultLogger.Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	defaultLogger.Infow(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	defaultLogger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	defaultLogger.Errorw(msg, keysAndValues...)
}
func Panic(msg string, keysAndValues ...interface{}) {
	defaultLogger.Panicw(msg, keysAndValues...)
}

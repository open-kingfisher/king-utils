package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugar *zap.SugaredLogger
var atom zap.AtomicLevel
var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func SetLoggerLevel(level string) {
	if l, ok := levelMap[level]; ok {
		atom.SetLevel(l)
	} else {
		atom.SetLevel(zapcore.InfoLevel)
	}
}

func init() {
	// 设置日志级别，默认info级别
	atom = zap.NewAtomicLevel()
	atom.SetLevel(zapcore.InfoLevel)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), os.Stdout, atom)
	logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),                  // caller 到真正的代码路径
		zap.AddStacktrace(zapcore.ErrorLevel), // 在error级别启动stacktrace
	)
	// Sync刷新任何缓冲的日志条目
	defer logger.Sync()
	sugar = logger.Sugar()
	zap.NewAtomicLevelAt(zapcore.FatalLevel)
}

func NewZap() *zap.Logger {
	return logger
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infof(format string, args ...interface{}) {
	sugar.Infof(format, args...)
}

func Infow(msg string, args ...interface{}) {
	sugar.Infow(msg, args...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	sugar.Warnf(format, args...)
}

func Warnw(msg string, args ...interface{}) {
	sugar.Warnw(msg, args...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}

func Errorw(msg string, args ...interface{}) {
	sugar.Errorw(msg, args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	sugar.Fatalf(format, args...)
}

func Fatalw(msg string, args ...interface{}) {
	sugar.Fatalw(msg, args...)
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	sugar.Debugf(format, args...)
}

func Debugw(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}

func DPanic(args ...interface{}) {
	sugar.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	sugar.DPanicf(format, args...)
}

func DPanicw(msg string, args ...interface{}) {
	sugar.DPanicw(msg, args...)
}

func Panic(args ...interface{}) {
	sugar.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	sugar.Panicf(format, args...)
}

func Panicw(msg string, args ...interface{}) {
	sugar.Panicw(msg, args...)
}

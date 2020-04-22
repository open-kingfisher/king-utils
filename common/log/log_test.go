package log

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

const input = "this is %s log ..."

func TestZap(t *testing.T) {
	SetLoggerLevel("debug")
	DPanicf(input, "dPanicf")
	Debugf(input, "debugf")
	Info("this is info log ...")
	Infof(input, "infof")
	Infow("this is infow log",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("uid", 3),
		zap.Duration("backoff", time.Second),
	)
	Warnf(input, "warnf")
	Errorf(input, "errorf")
	Fatalf(input, "fatalf")
	Panicf(input, "Panicf")
}

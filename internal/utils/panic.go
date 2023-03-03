package utils

import (
	"db-go-gin/internal/global"
	"runtime"

	"go.uber.org/zap"
)

func PrintPanic() {
	if err := recover(); err != nil {
		global.LOG.Error("Panic: %v", zap.Any("err", err))
		buf := make([]byte, 2048)
		n := runtime.Stack(buf, false)
		global.LOG.Error("Panic: %s", zap.String("str", string(buf[:n])))
	}
}

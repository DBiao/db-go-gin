package logger

import (
	"db-go-gin/internal/global"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

// GetWriteSyncer
// @description: zap logger中加入file-rotatelogs
// @return: zapcore.WriteSyncer, error
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join("./logs", "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(30*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

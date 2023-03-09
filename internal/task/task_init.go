package task

import (
	"time"
)

type Task interface {
	Start() bool
	Stop() bool
}

// Timer
/**
 * 定时调用
 * @delay 首次延时
 * @tick  间隔
 * @fun   定时执行function
 * @param fun参数
 */
func Timer(delay, tick time.Duration, task Task) {
	go func() {
		defer func() {
			task.Stop()
		}()

		t := time.NewTimer(delay)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				if task.Start() {
					return
				}
				if t.Reset(tick) {
					return
				}
			}
		}
	}()
}

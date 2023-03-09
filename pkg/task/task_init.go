package task

import (
	"time"
)

type TimerFunc func(interface{}) bool

type Task interface {
	start() bool
	stop() bool
}

// Timer
/**
 * 定时调用
 * @delay 首次延时
 * @tick  间隔
 * @fun   定时执行function
 * @param fun参数
 */
func timer(delay, tick time.Duration, task Task) {
	go func() {
		defer func() {
			task.stop()
		}()

		t := time.NewTimer(delay)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				if task.start() {
					return
				}
				if t.Reset(tick) {
					return
				}
			}
		}
	}()
}

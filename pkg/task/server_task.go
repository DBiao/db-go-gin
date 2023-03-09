package task

import (
	"fmt"
	"time"
)

type task1 struct {
}

func init() {
	task := &task1{}
	timer(2*time.Second, 5*time.Second, task)
}

// Start 注册
func (t *task1) start() bool {
	fmt.Println("start")
	return false
}

// stop 下线
func (t *task1) stop() bool {
	fmt.Println("stop")
	return true
}

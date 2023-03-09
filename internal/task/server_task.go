package task

import (
	"fmt"
	"time"
)

type Task1 struct{}

func init() {
	task := &Task1{}
	Timer(2*time.Second, 5*time.Second, task)
}

// Start 注册
func (t *Task1) Start() bool {
	fmt.Println("start")
	return false
}

// Stop 下线
func (t *Task1) Stop() bool {
	fmt.Println("stop")
	return true
}

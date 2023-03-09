package snowflake

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	if err := Init(1); err != nil {
		fmt.Printf("Init failed,err:%v\n", err)
		return
	}
	id, err := GetID()
	if err != nil {
		fmt.Printf("get Id failed,err:%v\n", err)
		return
	}
	fmt.Println(id)
}

func TestB(t *testing.T) {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetId())
	}
}

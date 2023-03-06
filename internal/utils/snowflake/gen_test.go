package snowflake

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
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

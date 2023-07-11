package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestAsyncEventBus_Publish(t *testing.T) {
	// 初始化消息总线
	bus := NewAsyncEventBus()
	// 订阅主题
	err := bus.Subscribe(
		"topic:1", func(msg1, msg2 string) {
			time.Sleep(1 * time.Microsecond)
			fmt.Printf("sub1, %s %s\n", msg1, msg2)
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	err = bus.Subscribe(
		"topic:1", func(msg1, msg2 string) {
			fmt.Printf("sub2, %s %s\n", msg1, msg2)
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	// 发布消息
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}

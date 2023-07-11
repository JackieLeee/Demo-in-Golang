package simplefactory

import "testing"

/**
 * @Author  Jackie
 * @Date  2020/8/22
 * @Description 简单工厂模式
 */

func TestHiAPI_Say(t *testing.T) {
	api := NewAPI(1)
	if api == nil {
		t.Error("NewAPI return nil")
	}
	if api.Say("Jackie") != "Hi, Jackie" {
		t.Error("Say test failed")
	}
}

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(2)
	if api == nil {
		t.Error("NewAPI return nil")
	}
	if api.Say("Jackie") != "Hello, Jackie" {
		t.Error("Say test failed")
	}
}

package singleton

import "testing"

/**
 * @Author  Jackie
 * @Date  2020/8/22
 * @Description 单例模式
 */

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Error("instance is not equal")
	}
}

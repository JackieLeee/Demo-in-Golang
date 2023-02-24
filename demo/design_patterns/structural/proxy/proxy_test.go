package proxy

import "testing"

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 代理模式
 */

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()

	if res != "pre:real:post" {
		t.Fail()
	}
}

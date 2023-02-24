package facade

import "testing"

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 外观模式
 */

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}

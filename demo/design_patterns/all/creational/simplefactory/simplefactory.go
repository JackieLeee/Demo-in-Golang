package simplefactory

import "fmt"

/**
 * @Author  Jackie
 * @Date  2020/8/22
 * @Description 简单工厂模式
 */

// API api接口
type API interface {
	Say(name string) string
}

// NewAPI 返回api实例
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI api接口的一个实现者
type hiAPI struct{}

// Say api接口say方法的实现
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI api接口的一个实现者
type helloAPI struct{}

// Say api接口say方法的实现
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

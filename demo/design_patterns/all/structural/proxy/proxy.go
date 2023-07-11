package proxy

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 代理模式
 */

// Subject 抽象主题角色
type Subject interface {
	Do() string
}

// RealSubject 具体主题角色
type RealSubject struct{}

func (RealSubject) Do() string {
	return ":real"
}

// Proxy 代理角色
type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等
	res += "pre"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等
	res += ":post"

	return res
}

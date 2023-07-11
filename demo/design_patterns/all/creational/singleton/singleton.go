package singleton

import "sync"

/**
 * @Author  Jackie
 * @Date  2020/8/22
 * @Description 单例模式
 */

// Singleton 单例模式的接口
type Singleton interface {
	foo()
}

// singleton 单例模式类，包私有的
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 获取单例模式对象
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

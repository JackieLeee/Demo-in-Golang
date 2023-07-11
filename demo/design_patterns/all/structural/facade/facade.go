package facade

import "fmt"

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 外观模式
 */

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API 接口
type API interface {
	Test() string
}

// apiImpl 具体实现
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

// Test 实现接口
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// NewAModuleAPI 返回一个aModule实现
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI aModule接口
type AModuleAPI interface {
	TestA() string
}

// aModuleImpl aModule实现
type aModuleImpl struct{}

// TestA 实现接口
func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI 返回一个bModule实现
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI bModule接口
type BModuleAPI interface {
	TestB() string
}

// bModuleImpl bModule实现
type bModuleImpl struct{}

// TestB 实现接口
func (*bModuleImpl) TestB() string {
	return "B module running"
}

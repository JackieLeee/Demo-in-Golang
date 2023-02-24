package decorator

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 装饰模式
 */

// Component 抽象构件角色
type Component interface {
	Calc() int
}

// ConcreteComponent 具体构件角色
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// MulDecorator 具体装饰角色
type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// AddDecorator 具体装饰角色
type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

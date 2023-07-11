package decorator

import "fmt"

/**
 * @Author: Jackie
 * @Date: 2020/8/22
 * @Description: 装饰模式
 */

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10) // 0 + 10
	c = WarpMulDecorator(c, 8)  // 10 * 8
	res := c.Calc()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}

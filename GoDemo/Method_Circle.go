package main

import "fmt"

//定义结构体
type Circle struct {
	radius float64
}

//该method属于Circle类型对象中的方法
func (c Circle) getArea () float64 {
	//c.radius 即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func (c Circle) getLenCircle () float64 {
	return 2 * 3.14 * c.radius
}

func main () {
	var c Circle
	c.radius = 10.00
	fmt.Printf("圆的面积 = %f\n", c.getArea())
	fmt.Printf("圆的周长 = %f\n", c.getLenCircle())
}
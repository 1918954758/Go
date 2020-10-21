package main

import "unsafe"

/*
 * 常量的表达式必须是使用内置函数，否则编译不通过
 * 内置函数，如：len()   cap()   unsafe()   Sizeof()
 */

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	d = unsafe.Sizeof(13)
)
func main() {
	println(a, b, c, d)
}

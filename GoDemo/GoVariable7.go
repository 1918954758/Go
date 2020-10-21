package main

import "fmt"

/*
 * 测试iota用法
 */
func main() {
	const (
		a = iota		// 0
		b				//1
		c				//2
		d = "ha"		//独立值， iota += 1
		e				//ha	iota += 1
		f = 100			//100   iota += 1
		g				//
		h = iota		//7  iota恢复计数
		i
	)
	const (
		ii = 1 << iota
		jj = 3 << iota
		kk
		ll
	)
	println(a, b, c, d, e, f, g, h ,i)
	println()
	fmt.Println("ii = ", ii)
	fmt.Println("jj = ", jj)
	fmt.Println("kk = ", kk)
	fmt.Println("ll = ", ll)
}

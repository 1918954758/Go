package main

import "fmt"

/*
 * 测试fallthrough，会报错，但是也可以执行
 */
func main () {
	switch {
		case false:
			fmt.Println("1. case条件语句为false")
			fallthrough
		case true:
			fmt.Println("2. case条件语句为true")
			fallthrough
		case false:
			fmt.Println("3. case条件语句为false")
			fallthrough
		case true:
			fmt.Println("4. case条件语句为true")
		case false:
			fmt.Println("5. case条件语句为false")
			fallthrough
		default:
			fmt.Println("6. 默认case")
	}
}

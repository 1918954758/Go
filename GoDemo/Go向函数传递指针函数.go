package main

import "fmt"

func main () {
	var a int = 100
	var b int = 200
	fmt.Print("\n---- 交换前 ---\n")
	fmt.Printf("%d\t%d", a, b)
	swap(&a, &b)
	fmt.Print("\n---- 交换后 ---\n")
	fmt.Printf("%d\t%d", a, b)
}
func swap (x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

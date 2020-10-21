package main

import "fmt"

/*
 * 求两个数的最大值
 */
func max (num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main () {
	result := max(5, 3)
	fmt.Println(result)
}

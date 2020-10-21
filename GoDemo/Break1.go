package main

import "fmt"

/* 当a大于15时，跳出循环 */
func main () {

	for a := 10; a < 20; a++ {
		fmt.Println(a)
		if a >= 15 {
			break
		}
	}
}

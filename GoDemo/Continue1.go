package main

import "fmt"

func main () {
	var a int = 10
	for a < 20 {
		if a ==15 {
			a++
			continue
		}
		fmt.Println(a)
		a++
	}
}

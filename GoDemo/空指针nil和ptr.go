package main

import "fmt"

func main () {
	var a int = 10
	var ptr *int
	ptr = &a
	fmt.Printf("ptr 的值为：%x\n", ptr)
	if ptr == nil {
		fmt.Println("ptr是空指针\n")
	}
	if ptr != nil {
		fmt.Println("ptr不是空指针\n")
	}
}


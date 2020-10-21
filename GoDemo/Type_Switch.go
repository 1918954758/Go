package main

import "fmt"

func main () {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x的类型是：%T", i)
	case float64:
		fmt.Printf("x的类型是：%T", i)
	case func(int) float64:
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Errorf("未知型")
	}
}

package main

import "fmt"

func main() {
	//定义了一个可以存储整数类型的带缓冲区通道
	//缓冲区大小为2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 6, -6, 7, 8, -8, -4, 0}
	c := make(chan int)
	go sum(v[:len(v)/2], c)
	go sum(v[len(v)/2:], c)
	x, y := <-c, <-c //从通道c中接收
	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, s := range s {
		sum += s
	}
	c <- sum //把sum发送到通道c
}

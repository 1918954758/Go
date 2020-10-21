package main

import "fmt"
func main () {
	slice := make([]int, 10, 20)    //定义slice的len=10，cap=20
	fmt.Printf("%p \n", slice) //slice指针地址
	fmt.Printf("%d %d\n", len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("%p \n", slice)  //没超过底层数组大小，地址不变
	fmt.Printf("%d %d\n", len(slice), cap(slice))

	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Printf("%p \n", slice)  //没超过底层数组大小，地址不变
	fmt.Printf("%d %d\n", len(slice), cap(slice))

	slice = append(slice, 11, 12, 13)
	fmt.Printf("%p \n", slice)  //超过底层数组大小，地址改变
	fmt.Printf("%d %d\n", len(slice), cap(slice))
}

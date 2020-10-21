package main

import "fmt"

func main () {
	var a int = 10//声明实际变量
	var ip *int//声明指针变量

	ip = &a//ip指针变量存储实际变量a的内存地址
	//a的内存地址
	fmt.Printf("a 的内存地址：%x\n", &a)
	//指针变量ip存储的变量a的内存地址
	fmt.Printf("指针变量ip存储的变量a的内存地址：%x\n", ip)
	//访问指针变量存储的内存地址的值
	fmt.Printf("访问指针变量存储的内存地址的值：%d\n", *ip)

}

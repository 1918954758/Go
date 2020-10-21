package main

import "fmt"

/*
 * 使用range循环来存取内存地址或值，是不可取的做法
 */
const MAX1 int = 3
func main () {
	number := [MAX1]int{5, 6, 7}
	var ptrs [MAX1]*int
	/* 使用range循环来存取内存地址或值，是不可取的做法 */
	/*
	不恰当写法
	for i, x := range &number {
		//将number数组的值得内存地址赋值给ptrs
		ptrs[i] = &x
	}
	*/
	/*
	正确做法1
	for i, x := range &number {
		temp := x
		ptrs[i] = &temp
	}
	 */
	for i := range number {
		ptrs[i] = &number[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组： 索引：%d  值：%d  值得内存地址：%d\n", i, *x, x)
	}

	/* 使用for循环来存取内存地址或值，是可取的做法 */
	for i := 0; i < MAX1; i++ {
		//将number数组的值得内存地址赋值给ptrs
		ptrs[i] = &number[i]
	}
	for i := 0; i < MAX1; i++ {
		fmt.Printf("指针数组： 索引：%d  值：%d  值得内存地址：%d\n", i, *ptrs[i], ptrs[i])
	}
}

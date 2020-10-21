package main

import "fmt"

const MAX int = 3
func main () {
	a := []int{10, 100, 200}
	var ptr [MAX]*int
	for i := 0; i < MAX; i++ {
		//将数组a中的分别对应存储到指针数组中
		ptr[i] = &a[i]
	}
	fmt.Println(a)
	fmt.Println(ptr)
	for j := 0; j < MAX; j++ {
		//取指针数组中的值
		fmt.Printf("*ptr[%d] = %d\t", j, *ptr[j])
	}
	fmt.Printf("\n\n")
	for k := 0; k < MAX; k++ {
		fmt.Printf("a[%d] = %d\t", k, a[k])
	}
}


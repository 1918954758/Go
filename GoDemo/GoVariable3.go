package main

import "fmt"

func main() {
	var i int//0
	var f32 float32//0
	var f64 float64//0
	var b bool//false
	var s string//""
	fmt.Printf("%v %v %v %v %q\n", i, f32, f64, b, s)

	var a1 *int//nil
	var a2 []int//[]
	var a3 map[string] int//map[]
	var a4 chan int//nil
	var a5 func(string) int//nil
	var a6 error   //nil     error是接口

	fmt.Println("\n", a1, a2, a3, a4, a5, a6)

	var t = true
	fmt.Println("\n", t)
}

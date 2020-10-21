package main

/*
 * &  *
 */
func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	println(b, c)

	ptr = &a
	println("a = ", a)
	println("ptr = ", ptr)
	println("&a = ", &a)
	println("&ptr = ", &ptr)
	println("*ptr = ", *ptr)
}

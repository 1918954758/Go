package main

import "fmt"

func main() {
	var i int = 7
	fmt.Println("i: ", &i)
	var j int = 7
	fmt.Println("j: ", &j)
	var k int = i
	fmt.Println("k: ", &k)
}

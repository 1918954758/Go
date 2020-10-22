package main

import "fmt"

//定义接口
type Ainterface interface {
	call()
}

//定义结构体1
type Astruct struct {
}

//定义结构体2
type Bstruct struct {
	name string
}

//入口main函数
func main() {
	var ainterface Ainterface
	ainterface = new(Astruct)
	ainterface.call()
	ainterface = new(Bstruct)
	ainterface.call()
}

//实现接口1
func (a Astruct) call() {
	fmt.Println("I am astruct, I can call you!")
}

//实现接口2
func (b Bstruct) call() {
	fmt.Println("I am bstruct, I can call you!")
}

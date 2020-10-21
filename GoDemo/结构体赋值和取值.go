package main
import "fmt"
type Bookss struct {
	title string
	author string
	subject string
	book_id int
}
func main () {
	var Book1 Bookss//声明Book1为Book类型
	var Book2 Bookss//声明Book2为Book类型

	//Book1描述
	Book1.title = "Go 语言"
	Book1.author = "www"
	Book1.subject = "Go"
	Book1.book_id = 12345

	//Book2描述
	Book2.title = "Go 语言"
	Book2.author = "wwf"
	Book2.subject = "Goo"
	Book2.book_id = 123456

	//打印Book1信息
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	//打印Book2信息
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

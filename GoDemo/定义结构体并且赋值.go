package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main () {
	fmt.Println(Books{"lemon", "lemon", "lemon", 12345})
	fmt.Println(Books{title:"lemon", author:"lemon", subject:"lemon", book_id:12345})
	fmt.Println(Books{title:"lemon", author:"lemon"})
}
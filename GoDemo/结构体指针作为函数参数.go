package main

import "fmt"

type Booked struct {
	title string
	author string
	subject string
	book_id int
}
func printBooked (book *Booked) {
	fmt.Printf("Bookl title : %s\n", book.title)
	fmt.Printf("Bookl author : %s\n", book.author)
	fmt.Printf("Bookl subjcet : %s\n", book.subject)
	fmt.Printf("Bookl book_id : %d\n", book.book_id)
}
func main () {
	var Book1 Booked
	var Book2 Booked

	Book1.title = "1"
	Book1.author = "2"
	Book1.subject = "3"
	Book1.book_id = 4

	Book2.title = "5"
	Book2.author = "6"
	Book2.subject = "7"
	Book2.book_id = 8

	printBooked(&Book1)
	printBooked(&Book2)
}

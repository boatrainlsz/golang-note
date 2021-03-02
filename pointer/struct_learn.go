package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	fmt.Println(Books{"go program", "bill", "gogo", 12})
	fmt.Println(Books{"go program", "bill", "gogo", 12})
	fmt.Println(Books{title: "go program", author: "bill", subject: "gogo", book_id: 12})
	fmt.Println(Books{title: "go program", author: "bill"})
	var book1 Books
	var book2 Books

	book1.title = "gdgd"
	book2.title = "dfdfdf"

	printBook(&book1)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

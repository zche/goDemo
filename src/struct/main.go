package main

import (
	"fmt"
	"struct/book"
)

func printBook(book book.Book) {
	fmt.Printf("id=%d,title=%s,author=%s,subject=%s\n",
		book.Id, book.Title, book.Author, book.Subject)
	book.Id = 100
}

func testBaseStruct() {
	var book1 *book.Book
	book1 = new(book.Book)
	book1.Id = 1001
	book1.Title = "go"
	book1.Author = "che"
	book1.Subject = "关于go的基本语法"
	fmt.Println(book1)
	book2 := book.Book{
		Id:      1002,
		Title:   "python",
		Author:  "che",
		Subject: "about python",
	}
	printBook(book2)
	book.RefTag(book2, 2)
	// fmt.Println(book1)
}

func testComplexStruct() {
	book.InitTechBook()
}

func main() {
	testBaseStruct()
	testComplexStruct()
}

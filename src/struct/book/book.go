package book

import (
	"fmt"
	"reflect"
)

type Book struct {
	Id      int    "这是id"
	Title   string "This is title"
	Author  string "This is author"
	Subject string "This is subject"
}

type techBook struct {
	cat string
	int
	Book
}

// NewBook的 构造函数
func NewBook(id int, title, author, subject string) *Book {
	return &Book{id, title, author, subject}
}

func (book *Book) String() string {
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n",
		book.Id, book.Title, book.Author, book.Subject)
}

func RefTag(b Book, idx int) {
	bType := reflect.TypeOf(b)
	idxField := bType.Field(idx)
	fmt.Printf("%v\n", idxField.Tag)
}

func InitTechBook() {
	bk := NewBook(1000, "C#", "jeffery", "about C#")
	tb := new(techBook)
	tb.cat = "tech"
	tb.int = 100
	tb.Book = *bk
	fmt.Println("techBook cat=", tb.cat)
	fmt.Println("techBook int=", tb.int)
	fmt.Println("techBook book=", tb.Book.String())
	fmt.Println("techBook=", tb.String())
}

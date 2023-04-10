package main

import (
	"fmt"
)

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b Book) DoNotSetPages(pages int) {
	b.pages = pages
}

type Bible struct {
	Book
	Type string
}

func (b Bible) Testament() string {
	return fmt.Sprintf("%s testament has %d pages\n", b.Type, b.Pages())
}

func (b *Bible) SetPages(pages int) {
	fmt.Printf("Child method has been called\n")
	b.Book.SetPages(pages)
}

func main() {
	var bigBook Book

	fmt.Printf("%v: %v\n", bigBook, bigBook.Pages())
	fmt.Printf("%T \n", bigBook.Pages)

	bigBook.SetPages(1500)
	bigBook.DoNotSetPages(1700)
	fmt.Println(bigBook.Pages())

	oldBook := Bible{Book{1500}, `old`}
	fmt.Println(oldBook.pages)

	oldBook.SetPages(1900)
	fmt.Println(oldBook.Testament())

	// exercise
	// 1. Make a roaring tiger, which extends a cat
	// implement: mewing and roaring
	// 2. Make a trash bin, which have an ability to accept garbage, can be filled or empty
}

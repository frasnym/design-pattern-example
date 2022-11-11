package main

import "fmt"

type BookCatalogAdapter struct {
	bookCatalog *Book
}

func (b *BookCatalogAdapter) printCatalogTitle() {
	fmt.Printf("%v by %v\n", b.bookCatalog.Title, b.bookCatalog.Author)
}

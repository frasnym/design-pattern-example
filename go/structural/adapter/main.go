package main

import "time"

func main() {
	catalogAdapter := &CatalogAdapter{}

	book := &Book{Title: "Java Book", Author: "Mr Java"}
	catalogAdapter.PrintCatalogTitle(&BookCatalogAdapter{bookCatalog: book})

	video := &Video{Title: "Java From 0", Author: "Ms Java", Duration: time.Hour * 1}
	catalogAdapter.PrintCatalogTitle(&VideoCatalogAdapter{videoCatalog: video})
}

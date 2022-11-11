package main

import "fmt"

type VideoCatalogAdapter struct {
	videoCatalog *Video
}

func (v *VideoCatalogAdapter) printCatalogTitle() {
	fmt.Printf("%v by %v\n", v.videoCatalog.Title, v.videoCatalog.Author)
}

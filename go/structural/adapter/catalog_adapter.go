package main

type CatalogAdapter struct {
}

func (*CatalogAdapter) PrintCatalogTitle(c ICatalogAdapter) {
	c.printCatalogTitle()
}

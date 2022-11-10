package main

type IStore interface {
	setName(newName string)
	setCity(newCity string)
	clone() IStore
	print()
}

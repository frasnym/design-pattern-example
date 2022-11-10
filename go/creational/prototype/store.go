package main

import "fmt"

type Store struct {
	Name     string
	City     string
	Country  string
	Category string
}

func (s *Store) clone() IStore {
	return s
}

func (s *Store) print() {
	fmt.Printf("%+v\n", s)
}

func (s *Store) setName(newName string) {
	s.Name = newName
}

func (s *Store) setCity(newCity string) {
	s.City = newCity
}

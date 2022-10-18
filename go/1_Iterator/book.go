package main

type Book struct {
	name string
}

func (b *Book) GetName() string {
	return b.name
}

func NewBook(name string) *Book {
	book := Book{name: name}
	return &book
}

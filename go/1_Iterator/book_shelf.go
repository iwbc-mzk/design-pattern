package main

type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf() *BookShelf {
	bs := &BookShelf{last: 0, books: []Book{}}
	return bs
}

func (b *BookShelf) GetBookAt(index int) Book {
	return b.books[index]
}

func (b *BookShelf) AppendBook(book Book) {
	b.books = append(b.books, book)
	b.last++
}

func (b *BookShelf) GetLength() int {
	return b.last
}

func (b *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(*b)
}

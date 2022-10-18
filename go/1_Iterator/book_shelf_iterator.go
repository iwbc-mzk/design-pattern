package main

type BookShelfIterator struct {
	bookShelf BookShelf
	index int
}

func NewBookShelfIterator(bookShelf BookShelf) *BookShelfIterator {
	b := &BookShelfIterator{bookShelf: bookShelf, index: 0}
	return b
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.bookShelf.GetBookAt(b.index)
	if b.index < b.bookShelf.GetLength() {
		b.index++
	}
	return book
}
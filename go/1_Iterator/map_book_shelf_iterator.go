package main

type MapBookShelfIterator struct {
	mapBookShelf MapBookShelf
	index int
}

func NewMapBookShelfIterator(bookShelf MapBookShelf) *MapBookShelfIterator {
	b := &MapBookShelfIterator{mapBookShelf: bookShelf, index: 0}
	return b
}

func (b *MapBookShelfIterator) HasNext() bool {
	if b.index < b.mapBookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (b *MapBookShelfIterator) Next() interface{} {
	book := b.mapBookShelf.GetBookAt(b.index)
	if b.index < b.mapBookShelf.GetLength() {
		b.index++
	}
	return book
}
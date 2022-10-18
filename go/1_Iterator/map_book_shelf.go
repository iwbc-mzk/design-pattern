package main

type MapBookShelf struct {
	books map[string]Book
	keys []string
	last  int
}

func NewMapBookShelf() *MapBookShelf {
	bs := &MapBookShelf{last: 0, books: map[string]Book{}}
	return bs
}

func (b *MapBookShelf) GetBookAt(index int) Book {
	key := b.keys[index]
	return b.books[key]
}

func (b *MapBookShelf) AppendBook(book Book) {
	name := book.GetName()
	b.books[name] = book
	b.keys = append(b.keys, name)
	b.last++
}

func (b *MapBookShelf) GetLength() int {
	return b.last
}

func (b *MapBookShelf) Iterator() Iterator {
	return NewMapBookShelfIterator(*b)
}

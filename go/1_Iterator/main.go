package main

import (
	"fmt"
)

func main() {
	bs := NewBookShelf()
	bs.AppendBook(*NewBook("Anime"))
	bs.AppendBook(*NewBook("Bike"))
	bs.AppendBook(*NewBook("Candy"))
	bs.AppendBook(*NewBook("Days"))

	// ここで利用しているのはIteratorオブジェクトのHasNext(), Next()のみ
	// → 集合体であるBookShelfの実装に依存せずに要素の走査を行っている
	it := bs.Iterator()
	for it.HasNext() {
		book := (it.Next()).(Book)
		fmt.Println(book.GetName())
	}

	mbs := NewMapBookShelf()
	mbs.AppendBook(*NewBook("Anime (Map)"))
	mbs.AppendBook(*NewBook("Bike (Map)"))
	mbs.AppendBook(*NewBook("Candy (Map)"))
	mbs.AppendBook(*NewBook("Days (Map)"))

	// MapBookShelfでもBookShelfと同様の方法で集合体の走査ができている
	mit := mbs.Iterator()
	for mit.HasNext() {
		book := (mit.Next()).(Book)
		fmt.Println(book.GetName())
	}
}
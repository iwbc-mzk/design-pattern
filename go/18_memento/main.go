package main

import (
	"fmt"
	"memento/memento"
)

func main() {

	game := memento.NewGame()
	caretaker := memento.NewCareTaker(game)

	game.Show()

	fmt.Println("== make backup ==")
	caretaker.MakeBackup()

	fmt.Println("== set mark at (1, 1) ==")
	game.SetMark(1, 1)

	fmt.Println("== make backup ==")
	caretaker.MakeBackup()

	game.Show()

	fmt.Println("== set mark at (3, 2) ==")
	game.SetMark(3, 2)

	game.Show()

	fmt.Println("== Undo ==")
	caretaker.Undo()

	game.Show()

	fmt.Println("== Undo ==")
	caretaker.Undo()

	game.Show()
}

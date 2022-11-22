package main

import "fmt"

func main() {
	editor := NewEditor()
	app := NewApp(editor)

	fmt.Println("=== コマンド実行 ===")

	app.Input("Hello World")
	fmt.Println()

	app.Input("This is Command Pattern")
	fmt.Println()

	app.Clear()
	fmt.Println()

	fmt.Println("=== Undo実行 ===")

	app.Undo()
	fmt.Println()

	app.Undo()
	fmt.Println()

	app.Undo()
	fmt.Println()
}
package main

func main() {
	cd := NewTemplateDisplay(NewCharDisplay('a'))
	sd := NewTemplateDisplay(NewStringDisplay("str"))

	cd.Display()
	sd.Display()

	cd.DisplayFiveTimes()
	sd.DisplayFiveTimes()
}
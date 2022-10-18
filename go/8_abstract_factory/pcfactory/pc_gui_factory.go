package pcfactory

import (
	. "abstract_factory/abstract_factory"
)

type pcGuiFactory struct {}

func NewGetPcGuiFactory() *pcGuiFactory {
	return &pcGuiFactory{}
}

func (f *pcGuiFactory) CreateCheckBox() ICheckBox {
	return &pcCheckBox{
		CheckBox: CheckBox{Height: 20, Width: 20},
	}
}

func (f *pcGuiFactory) CreateButton() IButton {
	return &pcButton{
		Button: Button{Color: "red", Text: "PC Button"},
	}
}
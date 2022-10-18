package mobilefactory

import (
	. "abstract_factory/abstract_factory"
)

type mobileGuiFactory struct {}

func NewGetMobileGuiFactory() *mobileGuiFactory {
	return &mobileGuiFactory{}
}

func (f *mobileGuiFactory) CreateCheckBox() ICheckBox {
	return &mobileCheckBox{
		CheckBox: CheckBox{Height: 10, Width: 10},
	}
}

func (f *mobileGuiFactory) CreateButton() IButton {
	return &mobileButton{
		Button: Button{Color: "blue", Text: "Mobile Button"},
	}
}
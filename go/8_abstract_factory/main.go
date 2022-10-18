package main

import (
	"fmt"
	. "abstract_factory/abstract_factory"
	. "abstract_factory/pcfactory"
	. "abstract_factory/mobilefactory"
)

func getFactory(factoryType string) IGuiFactory {
	if factoryType == "pc" {
		return NewGetPcGuiFactory()
	}
	if factoryType == "mobile" {
		return NewGetMobileGuiFactory()
	}
	return nil
}

type app struct {
	factory IGuiFactory
}

func (a *app) createGUI() {
	c := a.factory.CreateCheckBox()
	b := a.factory.CreateButton()
	fmt.Printf("CheckBox={height: %d, width: %d}\n", c.GetHeight(), c.GetWidth())
	fmt.Printf("Button={color: %s, text: %s}\n", b.GetColor(), b.GetText())
}

func main() {
	types := []string{"pc", "mobile"}

	// クライアントからは抽象型(IGuiFactory, ICheckBox, IButton)を通して処理を行う。
	// これにより仮に新たな型・ファクトリが追加されてもクライアント側は問題なく動作する。
	for _, t := range types {
		fmt.Println(t)
		f := getFactory(t)
		a := app{factory: f}
		a.createGUI()
	}
}
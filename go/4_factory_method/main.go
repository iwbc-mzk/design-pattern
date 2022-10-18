package main

import (
	. "factory_method/framework"
	. "factory_method/idcard"
	. "factory_method/cellphone"
	"fmt"
)

func main() {
	// ID Card 生成
	idf := NewIDCardFactory()
	f := NewTemplateFactory(idf)

	c1 := f.Create("田中太郎")
	c2 := f.Create("名無しの権兵衛")
	c3 := f.Create("山田花子")

	c1.Use()
	c2.Use()
	c3.Use()

	fmt.Println(idf.GetOwners())

	fmt.Println()

	// Cell Phone 生成
	cpf := NewCellPhoneFactory()
	cf := NewTemplateFactory(cpf)

	cf1 := cf.Create("田中太郎")
	cf2 := cf.Create("名無しの権兵衛")
	cf3 := cf.Create("山田花子")

	cf1.Use()
	cf2.Use()
	cf3.Use()

	fmt.Println(cpf.GetOwners())
}

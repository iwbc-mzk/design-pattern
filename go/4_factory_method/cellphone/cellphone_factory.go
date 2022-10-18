package cellphone

import (
	. "factory_method/framework"
)

type CellPhoneFactory struct {
	owners map[int]string
	serial int
}

func NewCellPhoneFactory() *CellPhoneFactory {
	f := CellPhoneFactory{owners: map[int]string{}, serial: 0}
	return &f
}

func (i *CellPhoneFactory) CreateProduct(owner string) Product {
	return newCellPhone(owner)
}

func (i *CellPhoneFactory) RegisterProduct(product *Product) {
	p := (*product).(CellPhone)
	i.owners[i.serial] = p.GetOwner()
	i.serial++
}

func (i *CellPhoneFactory) GetOwners() map[int]string {
	return i.owners
}

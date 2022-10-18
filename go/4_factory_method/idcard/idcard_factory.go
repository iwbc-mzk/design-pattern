package idcard

import (
	. "factory_method/framework"
)

type IDCardFactory struct {
	owners []string
}

func NewIDCardFactory() *IDCardFactory {
	f := IDCardFactory{owners: []string{}}
	return &f
}

func (i *IDCardFactory) CreateProduct(owner string) Product {
	return newIDCard(owner)
}

func (i *IDCardFactory) RegisterProduct(product *Product) {
	p := (*product).(IDCard)
	i.owners = append(i.owners, p.GetOwner())
}

func (i *IDCardFactory) GetOwners() []string {
	return i.owners
}

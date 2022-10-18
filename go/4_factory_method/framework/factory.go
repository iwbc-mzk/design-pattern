package framework

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product *Product)
}

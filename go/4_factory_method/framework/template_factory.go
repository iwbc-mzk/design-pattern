package framework

// 構造体作成のテンプレートを定義している構造体
// Factoryインターフェースを実装した構造体を利用してProduct構造体を作成する
type TemplateFactory struct {
	factory Factory
}

func NewTemplateFactory(factory Factory) *TemplateFactory {
	f := TemplateFactory{factory: factory}
	return &f
}

// Factory Method (Template Method)
func (t *TemplateFactory) Create(owner string) Product {
	p := t.factory.CreateProduct(owner)
	t.factory.RegisterProduct(&p)
	return p
}

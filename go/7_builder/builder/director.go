package builder

// BuilderのAPIを用いてインスタンスを生成する
// Builderの具象クラスには依存しない
type Director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *Director) Build() interface{} {
	d.builder.SetHandle()
	d.builder.SetTireSize()
	d.builder.SetFrame()
	// 最終的なオブジェクトはビルダーで生成したものを返す
	return d.builder.Build()
}
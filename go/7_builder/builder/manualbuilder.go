package builder

type ManualBuilder struct {
	manual
}

func NewManualBuilder() *ManualBuilder {
	return &ManualBuilder{}
}

func (b *ManualBuilder) SetHandle() {
	b.handlePage = "handel manual"
}

func (b *ManualBuilder) SetTireSize() {
	b.tirePage = "tire manual"
}

func (b *ManualBuilder) SetFrame() {
	b.framePage = "frame manual"
}

func (b *ManualBuilder) Build() interface{} {
	return manual{
		handlePage: b.handlePage,
		tirePage: b.tirePage,
		framePage: b.framePage,
	}
}
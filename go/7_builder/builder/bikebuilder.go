package builder

type BikeBuilder struct {
	bike
}

func NewBikeBuilder() *BikeBuilder {
	return &BikeBuilder{}
}

func (b *BikeBuilder) SetHandle() {
	b.handle = "bike handle"
}

func (b *BikeBuilder) SetTireSize() {
	b.tireSize = 30
}

func (b *BikeBuilder) SetFrame() {
	b.frame = "bike frame"
}

func (b *BikeBuilder) Build() interface{} {
	return bike{
		handle: b.handle,
		tireSize: b.tireSize,
		frame: b.frame,
	}
}
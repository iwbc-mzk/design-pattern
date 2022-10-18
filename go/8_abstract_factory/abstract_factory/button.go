package abstract_factory

type IButton interface {
	SetColor(color string)
	SetText(text string)
	GetColor() string
	GetText() string
}

type Button struct {
	Color string
	Text string
}

func (b *Button) SetColor(color string) {
	b.Color = color
}

func (b *Button) SetText(text string) {
	b.Text = text
}

func (b *Button) GetColor() string {
	return b.Color
}

func (b *Button) GetText() string {
	return b.Text
}
package abstract_factory

type ICheckBox interface {
	SetHeight(height int)
	SetWidth(width int)
	GetHeight() int
	GetWidth() int
}

type CheckBox struct {
	Height int
	Width int
}

func (c *CheckBox) SetHeight(height int) {
	c.Height = height
}

func (c *CheckBox) SetWidth(width int) {
	c.Width = width
}

func (c *CheckBox) GetHeight() int {
	return c.Height
}

func (c *CheckBox) GetWidth() int {
	return c.Width
}
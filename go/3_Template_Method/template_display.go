package main

type TemplateDisplay struct {
	display Display
}

func NewTemplateDisplay(d Display) *TemplateDisplay{
	t := TemplateDisplay{display: d}
	return &t
}

func (d *TemplateDisplay) Display() {
	d.display.Open()
	d.display.Print()
	d.display.Close()
}

func (d *TemplateDisplay) DisplayFiveTimes() {
	d.display.Open()
	for i := 0; i < 5; i++ {		
		d.display.Print()
	}
	d.display.Close()
}
package main

// Adapter
type PrintBanner struct {
	banner Banner
}

func NewPrintBanner(str string) *PrintBanner {
	b := NewBanner(str)
	pb := PrintBanner{banner: *b}
	return &pb
}

func (p PrintBanner) PrintWeak() {
	p.banner.ShowWithParen()
}

func (p PrintBanner) PrintStrong() {
	p.banner.ShowWithAster()
}

package main

func main() {
	p := NewPrintBanner("Hello")
	// PrintインターフェースのPrintWeak, PrintStrongでそれぞれBannerのShowWithParen, ShowWithAsterを実行できている。
	p.PrintWeak()
	p.PrintStrong()
}

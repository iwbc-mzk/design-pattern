package main

import (
	"fmt"
	"facade/counter"
)

func main() {
	// 利用する側はFacadeの背後にある複雑なサブシステムに依存しない。
	// サブシステムの修正や別サブシステムへの切替の場合は、Facadeクラスの変更のみが必要になる。
	wc := counter.NewWordCounter(" ")
	count := wc.Count("This is a pen. This pen is so good. I like this pen.")
	fmt.Println(count)
}
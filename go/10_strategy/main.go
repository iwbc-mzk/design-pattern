package main

import (
	"fmt"
)

func main() {
	c := NewContext()
	ed := NewDollarExchanger()
	ee := NewEuroExchanger()

	yenList := []float64{100.0, 145.3, 1009.3}

	c.SetExchanger(ed)
	fmt.Println("円 -> ドル")
	for _, y := range yenList {
		fmt.Printf("%.3f円 -> %.3fドル\n", y, c.Exchange(y))
	}

	c.SetExchanger(ee)
	fmt.Println("円 -> ユーロ")
	for _, y := range yenList {
		fmt.Printf("%.3f円 -> %.3fユーロ\n", y, c.Exchange(y))
	}
}
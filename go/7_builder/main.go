package main

import (
	"fmt"
	. "builder/builder"
)

func main() {
	bb := GetBuilder("bike")
	d := NewDirector(bb)
	bike := d.Build()
	fmt.Printf("Bike: %+v\n", bike)

	mb := GetBuilder("manual")
	d.SetBuilder(mb)
	manual := d.Build()
	fmt.Printf("Manual: %+v\n", manual)
}
package main

import (
	"fmt"
)

func getOS(name string) OS {
	if name == "windows" {
		return NewWindows()
	}
	if name == "mac" {
		return NewMac()
	}
	return nil
}

func main() {
	ol := []string{"windows", "mac"}
	for _, o := range ol {
		os := getOS(o)

		fmt.Printf("\n<< %s - Application >>\n", o)
		app := NewApplication(os)
		app.Analize()

		fmt.Printf("\n<< %s - AdvancedApplication >>\n", o)
		ad := NewAdvancedApplication(os)
		ad.Analize()
		ad.ReleaseMemory()
	}
}
package main

import (
	"fmt"
)

// Concrete Implementations
type Windows struct {}

func NewWindows() *Windows {
	return &Windows{}
}

func (w *Windows) GetOS() string {
	return "Windows"
}

func (w *Windows) ScreenShot() {
	fmt.Println("[Windows] take screen shot")
}

func (w *Windows) AnalizeMemory() map[string]int {
	return map[string]int{"app1": 20, "app2": 30, "app3": 50}
}

func (w *Windows) AnalizeVolume() map[string]int {
	return map[string]int{"file1": 50, "file2": 10, "file3": 15}
}

func (w *Windows) Terminate(appName string) {
	fmt.Printf("[Windows] terminate %s\n", appName)
}

func (w *Windows) Delete(filePath string) {
	fmt.Printf("[Windows] delete %s\n", filePath)
}
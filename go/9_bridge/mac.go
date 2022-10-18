package main

import (
	"fmt"
)

// Concrete Implementation
type Mac struct {}

func NewMac() *Mac {
	return &Mac{}
}

func (m *Mac) GetOS() string {
	return "Mac"
}

func (m *Mac) ScreenShot() {
	fmt.Println("[Mac] take screen shot")
}

func (m *Mac) AnalizeMemory() map[string]int {
	return map[string]int{"app1": 90, "app2": 3, "app3": 7}
}

func (m *Mac) AnalizeVolume() map[string]int {
	return map[string]int{"file1": 10, "file2": 70, "file3": 20}
}

func (m *Mac) Terminate(appName string) {
	fmt.Printf("[Mac] terminate %s\n", appName)
}

func (m *Mac) Delete(filePath string) {
	fmt.Printf("[Mac] delete %s\n", filePath)
}
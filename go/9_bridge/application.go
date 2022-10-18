package main

import (
	"fmt"
)

// Abstraction
type Application struct {
	os OS
}

func NewApplication(os OS) *Application {
	return &Application{os: os}
}

func (a *Application) TakeScreenShot() {
	a.os.ScreenShot()
}

func (a *Application) Analize() {
	n := a.os.GetOS()
	m := a.os.AnalizeMemory()
	v := a.os.AnalizeVolume()
	fmt.Printf("OS: %s\nMemory: %v\nVolume: %v\n", n, m, v)
}

func (a *Application) Terminate(appName string) {
	a.os.Terminate(appName)
}

func (a *Application) Delete(filePath string) {
	a.os.Delete(filePath)
}
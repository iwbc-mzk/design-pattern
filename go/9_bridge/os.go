package main

// Implementation
type OS interface {
	GetOS() string
	ScreenShot()
	AnalizeMemory() map[string]int
	AnalizeVolume() map[string]int
	Terminate(appName string)
	Delete(filePath string)
}
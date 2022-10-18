package main

// Refined Abstraction
type AdvancedApplication struct {
	Application
}

func NewAdvancedApplication(os OS) *AdvancedApplication {
	return &AdvancedApplication{
		Application: *NewApplication(os),
	}
}

func (a *AdvancedApplication) ReleaseMemory() {
	m := a.os.AnalizeMemory()
	for k := range m {
		a.os.Terminate(k)
	}
}
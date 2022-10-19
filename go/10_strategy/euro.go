package main

type EuroExchanger struct {
	rate float64
}

func NewEuroExchanger() *EuroExchanger {
	return &EuroExchanger{rate: 0.00899}
}

func (e *EuroExchanger) Exchange(yen float64) float64 {
	return e.rate * yen
}

func (e *EuroExchanger) GetRate() float64 {
	return e.rate
}
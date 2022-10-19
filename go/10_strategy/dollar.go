package main

type DollarExchanger struct {
	rate float64
}

func NewDollarExchanger() *DollarExchanger {
	return &DollarExchanger{rate: 0.0193}
}

func (d *DollarExchanger) Exchange(yen float64) float64 {
	return d.rate * yen
}

func (d *DollarExchanger) GetRate() float64 {
	return d.rate
}
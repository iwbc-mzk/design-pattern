package main

type Exchanger interface {
	Exchange(yen float64) float64
	GetRate() float64
}
package main

import "fmt"

type IdValidator struct {}

func NewIdValidator() *IdValidator {
	return &IdValidator{}
}

func (v *IdValidator) IsValid(d *InputData) bool {
	return d.Id > 100
}

func (v *IdValidator) Done() {
	fmt.Println("'ID' is OK.")
}

func (v *IdValidator) Fail() {
	fmt.Println("'ID' is NG.")
}
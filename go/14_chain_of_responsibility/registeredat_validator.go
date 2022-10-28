package main

import (
	"fmt"
	"time"
)

type RegisteredAtValidator struct {}

func NewRegisteredAtValidator() *RegisteredAtValidator {
	return &RegisteredAtValidator{}
}

func (v *RegisteredAtValidator) IsValid(d *InputData) bool {
	limit := time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC)
	return d.RegisteredAt.After(limit)
}

func (v *RegisteredAtValidator) Done() {
	fmt.Println("'RegisteredAt' is OK.")
}

func (v *RegisteredAtValidator) Fail() {
	fmt.Println("'RegisteredAt' is NG.")
}
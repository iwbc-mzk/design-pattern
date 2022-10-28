package main

import (
	"time"
	"fmt"
)

func main() {
	// OK
	data1 := NewInputData(200, time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC), "Password1")
	// IDがNG
	data2 := NewInputData(10, time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC), "Password1")
	// 登録日がNG
	data3 := NewInputData(200, time.Date(2021, time.October, 1, 0, 0, 0, 0, time.UTC), "Password1")
	// パスワードがNG
	data4 := NewInputData(200, time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC), "password")

	iv := NewValidator(NewIdValidator())
	rv := NewValidator(NewRegisteredAtValidator())
	pv := NewValidator(NewPasswordValidator())
	iv.SetNext(rv).SetNext(pv)

	fmt.Println("data1 validation")
	fmt.Println("Validation result: ", iv.Validate(data1))

	fmt.Println()
	fmt.Println("data2 validation")
	fmt.Println("Validation result: ", iv.Validate(data2))

	fmt.Println()
	fmt.Println("data3 validation")
	fmt.Println("Validation result: ", iv.Validate(data3))

	fmt.Println()
	fmt.Println("data4 validation")
	fmt.Println("Validation result: ", iv.Validate(data4))
}
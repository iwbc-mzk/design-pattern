package main

import (
	"fmt"
	"regexp"
)

type PasswordValidator struct {}

func NewPasswordValidator() *PasswordValidator {
	return &PasswordValidator{}
}

func (v *PasswordValidator) IsValid(d *InputData) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(d.Password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(d.Password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(d.Password)
	return hasUpper && hasLower && hasNumber
}

func (v *PasswordValidator) Done() {
	fmt.Println("'Password' is OK.")
}

func (v *PasswordValidator) Fail() {
	fmt.Println("'Password' is NG.")
}
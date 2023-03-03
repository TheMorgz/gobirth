package validator

import (
	"fmt"
	"strings"
	"time"
)

type Validator struct {
	InputErrors map[string]string
}

func (v *Validator) AddInputError(key, message string) {
	if v.InputErrors == nil {
		v.InputErrors = make(map[string]string)
	}

	if _, exists := v.InputErrors[key]; !exists {
		v.InputErrors[key] = message
	}
}

func (v *Validator) CheckInput(ok bool, key, message string) {
	if !ok {
		v.AddInputError(key, message)
	}
}

func ValidDate(value string) bool {
	_, err := time.Parse("2006/01/02", value)
	if err != nil {
		return err == nil
	}
	return true
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) IsBirthday(date time.Time) bool {
	return (date.Day() == time.Now().Day())
}

func (v *Validator) IsFebTwentyNine(date time.Time, birthYear int, currentYear int) string {
	if date.Day() != 29 && date.Month() != 2 {
		return ""
	}

	isLeap := func(year int) bool {
		return year%400 == 0 || (year%4 == 0 && year%100 != 0)
	}

	if isLeap(birthYear) && !isLeap(currentYear) && (date.Day() == 29 && date.Month() == 2) {
		fmt.Println()
		return "Your birthday is on the 28th since there is no Feb 29th on this non-leap year!"
	}

	return ""
}

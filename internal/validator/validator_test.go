package validator

import (
	"testing"
	"time"
)

func TestAddInputError(t *testing.T) {
	v := &Validator{}
	key := "Name"
	message := "Name is required"

	v.AddInputError(key, message)

	if v.InputErrors == nil {
		t.Errorf("InputErrors map is nil")
	}

	if v.InputErrors[key] != message {
		t.Errorf("InputError message not added to map")
	}
}

func TestCheckInput(t *testing.T) {
	v := &Validator{}
	key := "Surname"
	message := "Surname is required"

	v.CheckInput(false, key, message)

	if v.InputErrors == nil {
		t.Errorf("InputErrors map is nil")
	}

	if v.InputErrors[key] != message {
		t.Errorf("InputError message not added to map")
	}
}

func TestValidDate(t *testing.T) {
	validDate := ValidDate("2022/01/01")

	if !validDate {
		t.Errorf("Valid date is not valid")
	}

	invalidDate := ValidDate("2022/13/01")

	if invalidDate {
		t.Errorf("Invalid date is valid")
	}
}

func TestNotBlank(t *testing.T) {
	notBlank := NotBlank("test")

	if !notBlank {
		t.Errorf("Not blank string is not valid")
	}

	blank := NotBlank("")

	if blank {
		t.Errorf("Blank string is valid")
	}
}

func TestIsBirthday(t *testing.T) {
	v := &Validator{}
	now := time.Now()
	birthday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	isBirthday := v.IsBirthday(birthday)

	if !isBirthday {
		t.Errorf("Today is not a birthday")
	}

	notBirthday := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)

	isNotBirthday := v.IsBirthday(notBirthday)

	if isNotBirthday {
		t.Errorf("Tomorrow is a birthday")
	}
}

func TestIsFebTwentyNine(t *testing.T) {
	validator := &Validator{}
	birthday := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
	birthYear := 2000
	currentYear := 2000

	// Test Leap Year
	result := validator.IsFebTwentyNine(birthday, birthYear, currentYear)
	expected := ""
	if result != expected {
		t.Errorf("TestIsFebTwentyNine(%s) = %s, expected %s", birthday, result, expected)
	}

	// Test non-Leap Year
	currentYear = 2001
	result = validator.IsFebTwentyNine(birthday, birthYear, currentYear)
	expected = "Your birthday is on the 28th since there is no Feb 29th on this non-leap year!"
	if result != expected {
		t.Errorf("TestIsFebTwentyNine(%s) = %s, expected %s", birthday, result, expected)
	}

	// Test non-February, non-29th date
	birthday = time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC)
	result = validator.IsFebTwentyNine(birthday, birthYear, currentYear)
	expected = ""
	if result != expected {
		t.Errorf("TestIsFebTwentyNine(%s) = %s, expected %s", birthday, result, expected)
	}
}

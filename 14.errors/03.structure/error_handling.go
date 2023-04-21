package main

import (
	"errors"
	"fmt"
)

type invalidInputError struct {
	message string
}

func (e invalidInputError) Error() string {
	return e.message
}

func validateInput(input int) error {
	if input < 0 {
		return invalidInputError{"input cannot be negative"}
	}

	return nil
}

func main() {
	err := validateInput(-5)
	if err != nil {
		var inputError *invalidInputError
		if errors.As(err, &inputError) {
			fmt.Println("Error:", err, "Please provide a valid input.")
		} else {
			fmt.Println("Error:", err)
		}
	}
}

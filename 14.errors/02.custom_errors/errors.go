package main

import "fmt"

type divisionError struct {
	message string
}

func New(s string) divisionError {
	return divisionError{s}
}

func (e divisionError) Error() string {
	return e.message
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, New("division by zero is not allowed")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

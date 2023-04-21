package main

import (
	"errors"
	"fmt"
)

func connectToDatabase() error {
	return errors.New("database connection failed")
}

func initializeApp() error {
	err := connectToDatabase()
	if err != nil {
		return fmt.Errorf("could not initialize app: %w", err)
	}
	return nil
}

var ErrDatabaseConnection = errors.New("database connection failed")

func main() {
	err := initializeApp()
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Is(err, ErrDatabaseConnection) {
			fmt.Println("The error is related to database connection.")
		}
	}
}

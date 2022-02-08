//bcrypt is a password hashing function.
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	s := `123456` //password
	sb, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(sb))

	loginPassword := `password123`

	err = bcrypt.CompareHashAndPassword(sb, []byte(loginPassword))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("You are logged in")
}

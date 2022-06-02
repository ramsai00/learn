package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	book1 := Book{Title: "You Win", Author: "S@i"}

	byteArray, err := json.Marshal(book1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(byteArray)

	fmt.Println("type casted to string: ", string(byteArray))
}

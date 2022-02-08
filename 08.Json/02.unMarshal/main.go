package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `[{"First":"Ram","Last":"sai","Age":21},{"First":"satya","Last":"venkata","Age":21}]`
	byteStream := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", byteStream)

	var p []person

	err := json.Unmarshal(byteStream, &p)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range p {
		fmt.Println("\n Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

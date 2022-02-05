package main

import "fmt"

// map are the key value store functions'
// map are the unordered list

func main() {
	m := map[string]int{
		"ram":    21,
		"satya":  27,
		"harish": 22,
	}
	fmt.Println(m)
	fmt.Println(m["ram"])
	fmt.Println(m["sai"]) // if the value is not present in go it returns 0
	v, ok := m["sai"]     //It is called comma(,) idiom.
	fmt.Println(ok)
	if ok {
		fmt.Println("present ?", v)
	}

	m["sundar"] = 25 //adding value to the existing map
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("key :", k, "value: ", v)
	}

	delete(m, "ram")
	fmt.Println(m)

	delete(m, "hkdf") // not throws an error even if the value doesn't exist
	fmt.Println(m)
}

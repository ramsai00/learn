package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	s := []string{"ram", "sai", "akash", "venu"}

	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}

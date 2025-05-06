package main

import "fmt"

func main() {
	x := make([]int, 3, 10)
	fmt.Println(x)
	x = []int{1, 2, 3}
	fmt.Println(x)
	//x[4] = 2					// exceeding length
	//fmt.Println(x)
	x = append(x, 4)
	fmt.Println(x)

	// MultiDimensional Slice
	names := []string{"ram", "sai", "harish", "satya", "shashank"}
	fmt.Println(names)
	city := []string{"rjy", "kkd", "gnt", "vjy", "elu"}
	fmt.Println(city)
	multi := [][]string{names, city}
	fmt.Println(multi)
}

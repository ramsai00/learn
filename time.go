package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now()

	//time.Time to Unix Timestamp
	tUnix := tNow.Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	//Unix Timestamp to time.Time
	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)
	wg.Wait()
}

func helloworld(wg *sync.WaitGroup) {
	fmt.Println("Hello World!")
	wg.Done()
}

func goodbye(wg *sync.WaitGroup) {
	fmt.Println("Good Bye!")
	wg.Done()
}

//You can use WaitGroups to wait for multiple goroutines to finish.
//A WaitGroup blocks the execution of a function until its internal counter becomes 0.

// wg.Add(int): This method indicates the number of goroutines to wait.
// In the above code, I have provided 2 for 2 different goroutines. Hence the internal counter wait becomes 2.
// wg.Wait(): This method blocks the execution of code until the internal counter becomes 0.
// wg.Done(): This will reduce the internal counter value by 1.
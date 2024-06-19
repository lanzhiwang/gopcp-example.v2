package main

import "fmt"

func main() {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}

// $ ./main
// 3
// 3
// 2
// 2
// 3

// $ ./main
// 3
// 2
// 3
// 3
// 1

// $ ./main
// 3
// 2
// 2
// 1
// 2

// $ ./main
// 2
// 3
// 1
// 3
// 3

// $

package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}

// $ ./main
// Present time: 2024-06-19 16:31:14.164673 +0800 CST m=+0.000160453.
// Expiration time: 2024-06-19 16:31:16.165716 +0800 CST m=+2.001194326.
// Stop timer: false.
// $

package main

import (
	"fmt"
	"time"

	"github.com/Sandeep4/throttler"
)

func main() {
	fmt.Println("Throtteling Example")
	rate := throttler.Rate{4, 10}
	rateLimiter := throttler.NewWindowThrottler(rate)
	for i := 0; i < 5; i++ {
		if !rateLimiter.ThrottleKey("hi") {
			fmt.Printf("Passed %d, ", i)
		} else {
			fmt.Printf("Failed %d\n", i)
		}
	}
	time.Sleep(10 * time.Second)
	for i := 0; i < 5; i++ {
		if !rateLimiter.ThrottleKey("hi") {
			fmt.Printf("Passed %d, ", i)
		} else {
			fmt.Printf("Failed %d\n", i)
		}
	}
}

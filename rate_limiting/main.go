package main

import (
	"fmt"
	"time"
)

//Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
//Go elegantly supports rate limiting with goroutines, channels, and tickers.
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	//short burst of requests in rate limiting.
	//Begin with 3 burst of requests,then one by one.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(1 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("requests", req, time.Now())
	}
}

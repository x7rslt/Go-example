package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future -
// tickers are for when you want to do something repeatedly at regular intervals.
func main() {
	t1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-t1.C:
				fmt.Println("Hello Everyone.", t)
			}
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	t1.Stop()
	done <- true
	fmt.Println("Ticker done.")
}

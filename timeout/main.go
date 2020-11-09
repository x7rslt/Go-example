package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "The first opration."
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("First opration timeout.")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "The second opration."
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("The second opration timeout.")
	}
}

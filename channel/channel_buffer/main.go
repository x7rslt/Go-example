package main

import "fmt"

//channel buffering
func main() {
	msg := make(chan string, 2)

	msg <- "test1"
	msg <- "test2"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

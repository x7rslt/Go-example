package main

import (
	"fmt"
	"time"
)

// method1:
// func main() {
// 	msg := make(chan string)
// 	go func() {
// 		fmt.Println("a channel")
// 		msg <- "work done"
// 	}()

// 	fmt.Println("task start:")
// 	<-msg
// }

//method2:suggest use this.
func worker(done chan bool) {
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

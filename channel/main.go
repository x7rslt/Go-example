package main

import "fmt"

//通道连接goroutine之间的通信
func main() {
	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()

	messages := <-msg
	fmt.Println(messages)

}

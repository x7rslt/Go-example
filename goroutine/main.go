package main

import (
	"fmt"
	"time"
)

func prtest(n string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %d %s\n", i, n)
	}
}

func main() {
	prtest("first")

	go prtest("second")

	//go 调用匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("lehehe")

	time.Sleep(time.Second)
	fmt.Println("done")
}

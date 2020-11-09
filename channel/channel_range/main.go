package main

import "fmt"

//Use range iterate over values received from a channel.
func main() {
	c := make(chan string, 2)
	c <- "test1"
	c <- "test2"
	close(c)

	for i := range c {
		fmt.Println(i)
		//fmt.Println(<-c)
		fmt.Println(1)
	}

}

package main

import (
	"fmt"
	"os"
)

//Use os.Exit to exit with a given status

func main() {

	//defer will never execute.
	defer fmt.Println("!")

	os.Exit(3) //3 is status. os.Exit() non-zero status print.
}

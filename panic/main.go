package main

import "os"

//use panic fail fast on errors that shouldn't occur during normal opreation.

func main() {
	//a panic will abort the programm.
	panic("this is a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

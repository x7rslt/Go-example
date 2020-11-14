package main

import (
	"fmt"
	"os"
)

//Defer is used to ensure that a function call is performed
//later ina program's execution,usually for purposes of cleanup.
//defer is often used where e.g.ensure and finally would be used inother languages.

func main() {
	f := createFile("./defer.txt")

	//Immediately after getting afile object with createFile,we defer
	//the closing of that file with closeFIle.This will be executed at
	//the end of the enclosing function(main),after writeFIle has finished.
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
}

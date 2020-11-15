package main

import (
	"flag"
	"fmt"
)

//command-line flags are a common way to specify options for command-line programs.
//For example,in wc -l the -l is a command-line flag.
func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork ", false, "a bool")

	var svar string
	//declare an option that uses an existing var declared elsewhere in the programm.Note need pointer.
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//execute the command-line prase
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool", *boolPtr)
	fmt.Println("svar:", svar)
	//trailing positional arguments
	fmt.Println("tail:", flag.Args())
}

//Execute programm:
//go run main.go -word=opt -numb=7 -fork -svar=flag
//go run main.go -h

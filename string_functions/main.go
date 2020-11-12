package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

//Standard library's strings package provides many useful string-related functions.

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "e"))
	p("Join", s.Join([]string{"a", "b"}, "-"))
	p("Replace", s.Replace("foo", "o", "0", -1))
	p("Replace", s.Replace("foo", "o", "0", 1))
	p("Repeat", s.Repeat("a", 5))
	p("Split", s.Split("a-b-c-d", "-"))
	p("Tolower", s.ToLower("TEST"))
	p("TOupper", s.ToUpper("test"))
	p()

	//Get length of a string in bytes and getting a byte by index.
	p("Len", len("hello"))
	p("Char", "Hello"[1])

}

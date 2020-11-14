package main

import (
	"fmt"
	"sort"
)

//Go’s sort package
//implements sorting for builtins and user-defined types.

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)

	//Note that sorting is in-place,
	//so it changes the given slice and doesn’t return a new one.
	fmt.Println(strs)

	ints := []int{1, 3, 4, 6}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}

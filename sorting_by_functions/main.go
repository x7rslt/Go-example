package main

import (
	"fmt"
	"sort"
)

//Custom sorts:For example, suppose we wanted to
//sort strings by their length instead of alphabetically.

// implement sort.Interface - Len, Less, and Swap - on our
//type so we can use the sort package’s generic Sort function.
type byLenth []string

func (s byLenth) Len() int {
	return len(s)
}

func (s byLenth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLenth) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLenth(fruits))
	fmt.Println(fruits)

}

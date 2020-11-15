package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseUint("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.Atoi("124")
	fmt.Println(u)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

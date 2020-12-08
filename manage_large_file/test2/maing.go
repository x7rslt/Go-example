package main

//so slow
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	find := []byte(os.Args[1])
	dat, err := ioutil.ReadFile(`E:\SocialData\person_data\weibo.txt`)
	check(err)
	if bytes.Contains(dat, find) {
		fmt.Print("yes")
	} else {
		fmt.Print("Not Found!")
	}
}

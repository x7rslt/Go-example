package main

//纯读取时间，未实际应用

import (
	"fmt"
	"io/ioutil"
	"time"
)

var m string = `E:\SocialData\person_data\weibo.txt`

func readAll(filePath string) {
	start := time.Now()
	ioutil.ReadFile(filePath)
	fmt.Println("read spend:", time.Now().Sub(start))
}

func main() {
	readAll(m)
}

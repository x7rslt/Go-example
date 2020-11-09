package main

import (
	"fmt"
)

//闭包:Go 支持匿名函数，实现闭包
func Add(m int) func(n int) int {
	return func(n int) int {
		num := n * m //运算要和后面的return拆分
		return num
	}
}

func main() {
	num := Add(10)
	fmt.Println(num(10))

}

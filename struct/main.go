//Struct 是一种数据类型的集合

package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person { //person是数据类型，获取值用：&person,返回类型：*person
	p := person{name: name}
	p.age = 18
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 18})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 19})
	fmt.Println(&person{"Ann", 19})
	fmt.Println(newPerson("John"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	fmt.Println(*sp)
	sp.age = 52
	fmt.Println(sp.age)
	h := newPerson("Tom")
	fmt.Println(*h)
}

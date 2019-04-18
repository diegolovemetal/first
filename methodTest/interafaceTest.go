package main

import (
	"fmt"
	"strconv"
)

type human struct {
	name string
	age int
	phone string
}

type student struct {
	human
	school string
	loan float32
}

type employee struct {
	human
	company string
	money float32
}

//human sayHi()

func (h human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//human Sing
func (h human) Sing(lyrics string) {
	fmt.Println("lalalala...", lyrics)
}

//employee sayHi()
func (e employee) sayHi() {
	fmt.Printf("Hi,I am %s,working at %s, Call me on %s", e.name, e.company, e.phone)
}
// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type men interface {
	sayHi()
	Sing(lyrics string)
}
//通过该方法human，实现了fmt.Stringer
func (h human) String() string {
	return "❰"+ h.name + " - " + strconv.Itoa(h.age) + "years - ✆" + h.phone + "❱"
}

func main() {
	//mike := Student{Human{"mike", 21, "12321312"}, "cqupt", 3}
	mike := student{human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := student{human{"Paul", 21, "213231244"},"CQUPT", 32000}
	sam := employee{human{"小刘", 25, "213124124"}, "Tecent", 12000}
	tom := employee{human{"老张", 30, "2131231234"}, "Bitdance", 20000}

	Bob := human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
	var i men
//i能存储student
	i = mike
	i.sayHi()
	i.Sing("If u like gamble")
	//i能储存employee
	i = tom
	i.sayHi()
	i.Sing("Born to be wild")

	//定义slice men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]men, 3)

	x[0], x[1], x[2] = mike, paul, sam

	for _, v :=  range x {
		v.sayHi()
	}

}

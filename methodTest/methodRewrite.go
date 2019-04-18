package main

import "fmt"

type Human struct {
	Name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employ struct {
	Human
	company string
	phone string
}

//Human 定义Method
func (h *Human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.phone)
}

func (e *Employ) sayHi()  {
	fmt.Printf("Hi,I am %s working at %s, my working phone is %s\n", e.Name, e.company, e.phone)
}

func main() {
	diego := Student{Human{"Diego", 22, "721231"}, "CQUPT"}
	sasa := Employ{Human{"Sasa", 21, "1231231"}, "CQUT", "023231231"}

	diego.sayHi()
	sasa.sayHi()
}
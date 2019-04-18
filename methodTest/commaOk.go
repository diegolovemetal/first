package main

import (
	"fmt"
	"strconv"
)

type Element interface {

}

type List []Element

type Person struct {
	name string
	age int
}

//定义了String方法，实现了fmt.Stringer

func (p Person) String() string {
	return "<name" + p.name + " - age " + strconv.Itoa(p.age) + " years>"
}

func main()  {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hi"
	list[2] = Person{"diego", 21}
//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，
// element是interface变量，T是断言的类型。
//	for index,element := range list {
//		if value, ok := element.(int); ok {
//			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
//		} else if value, ok := element.(string); ok {
//			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
//		} else if value, ok := element.(Person); ok {
//			fmt.Printf("list[%d] is a Person and its value is %s\n",index, value)
//		} else {
//			fmt.Printf("list[%d] is of a different type \n", index)
//		}
//	}

//除commaOk以外可以用switch来获取interface的类型
	for index, element := range list {
		switch value := element.(type) {
		//`element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)

		}
	}
}

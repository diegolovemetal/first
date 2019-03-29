package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	A string
	B int
}

func main() {
	t := Student{"diego", 123}
	s := reflect.ValueOf(&t).Elem()
	typeOfS := s.Type()

	for i:=0; i<s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfS.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetString("caonima")
	s.Field(1).SetInt(12)
	fmt.Println("t is now:",t)
}


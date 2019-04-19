package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 5.4
	v := reflect.ValueOf(x)
	fmt.Println("type", v.Type())
	fmt.Println("kind is float64", v.Kind() == reflect.Float64)
	fmt.Println("value", v.Float())
	y := reflect.ValueOf(&x)
	v = y.Elem()
	v.SetFloat(23.90)
	fmt.Println("value", v.Float())

}

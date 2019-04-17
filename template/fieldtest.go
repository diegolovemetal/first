package main

import (
	"html/template"
	"os"
)

type person struct {
	UserName string
}

func main()  {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}")
	p := person{UserName: "diego"}
	t.Execute(os.Stdout, p)
}
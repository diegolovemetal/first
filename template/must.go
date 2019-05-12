package main

import (
	"fmt"
	"html/template"
)

func main() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text	/* add a comment */"))
	fmt.Println("the first one parse ok!")

	template.Must(template.New("second").Parse("some static text {{.Name}}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse("some static text{{.Name}"))
}

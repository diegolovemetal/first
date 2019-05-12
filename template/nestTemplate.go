package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	s1, _ := template.ParseFiles("src/template/nestTemplate.go", "src/template/footer.tmpl", "src/template/content.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()

	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()

	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

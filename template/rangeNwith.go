package main

import (
	"html/template"
	"os"
)
type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails []string
	Friends []*Friend
}

func main() {
	f1 := Friend{Fname:"sasa"}
	f2 := Friend{Fname:"sunday"}

	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend's name is {{.Fname}}
		{{end}}
		{{end}}
		`)
	p := Person{UserName: "diego",
		Emails: []string{"123@123.com", "234@23.com"},
		Friends: []*Friend{&f1, &f2}}
		t.Execute(os.Stdout, p)

}

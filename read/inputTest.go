package main

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "17.1210 / 981021 / GO "
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("please enter your name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi,%s %s", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from Z string we read:", f, i, s )
}

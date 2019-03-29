package main

import (
	"bufio"
	"fmt"
	"os"
)

//var inputReader = *bufio.NewReader(os.Stdin)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input something:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("There was errors reading,exiting...")
		return
	}
	fmt.Printf("u input :%s\n",input)

	switch input {
	case "Diego\n":
		fmt.Println("welcome Diego!")
	default:
		fmt.Println("U r not welcome here.")
		}

}
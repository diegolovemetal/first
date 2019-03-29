package main

import (
"flag"
"fmt"
)

func main() {
	wordPtr := flag.String("word", "diego", "a string")
	numPtr := flag.Int("number", 21, "an int")
	boolPtr := flag.Bool("Fork", true, "a bool")
	var svar string
	flag.StringVar(&svar,"svar", "diegolovemetal", "a string var")

	//解析

	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

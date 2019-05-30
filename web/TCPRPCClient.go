package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A,B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "Server:Port")
		os.Exit(1)
	}

	service := os.Args[1]
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{13, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d / %d = %d\n Remainder = %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

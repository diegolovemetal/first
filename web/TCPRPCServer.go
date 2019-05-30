package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Arith int

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkErr(err error)  {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}
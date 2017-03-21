package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/netqyq/rpc/shared" //Path to the package contains shared struct
)

type Arith struct {
	client *rpc.Client
}

func (t *Arith) Divide(a, b int) shared.Quotient {
	args := &shared.Args{a, b}
	var reply shared.Quotient
	err := t.client.Call("Arithmetic.Divide", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func (t *Arith) Multiply(a, b int) int {
	args := &shared.Args{a, b}
	var reply int
	err := t.client.Call("Arithmetic.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func main() {

	// Tries to connect to localhost:1234 using HTTP protocol (The port on which rpc server is listening)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Create a struct, that mimics all methods provided by interface.
	// It is not compulsory, we are doing it here, just to simulate a traditional method call.
	arith := &Arith{client: client}

	fmt.Println(arith.Multiply(5, 6))
	fmt.Println(arith.Divide(500, 10))
}

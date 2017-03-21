package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/netqyq/rpc/shared" //Path to the package contains shared struct
)

// Every method that we want to export must have
// (1) the method has two arguments, both exported (or builtin) types
// (2) the method's second argument is a pointer
// (3) the method has return type error

type Arith int

func (t *Arith) Multiply(args *shared.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *shared.Args, quo *shared.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func registerArith(server *rpc.Server, arith shared.Arith) {
	// registers Arith interface by name of `Arithmetic`.
	// If you want this name to be same as the type name, you
	// can use server.Register instead.
	server.RegisterName("Arithmetic", arith)
}

func main() {

	//Creating an instance of struct which implement Arith interface
	arith := new(Arith)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Arith"
	// The wrapper method here ensures that only structs which implement Arith interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	registerArith(server, arith)

	// registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath
	server.HandleHTTP("/", "/debug")

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement starts go's http server on
	// socket specified by l.
	http.Serve(l, nil)
}

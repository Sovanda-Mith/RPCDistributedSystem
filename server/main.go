package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Define a struct. This struct will bind all the RPC methods
type Calculator struct{}

// Arguments for Multiply and Divide method
type Args struct {
	A, B int
}

// Multiply is the method that the client will call
func (c *Calculator) Multiply(args *Args, reply *int) error {
	if args == nil {
		return errors.New("arguments cannot be nil")
	}
	fmt.Printf("Arguments received: %d * %d\n", args.A, args.B)
	*reply = args.A * args.B
	fmt.Printf("Sending response to client: %d\n", *reply)
	return nil
}

// Divide is the method that the client will call
func (c *Calculator) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	fmt.Printf("Arguments received: %d / %d\n", args.A, args.B)
	*reply = args.A / args.B
	fmt.Printf("Sending response to client: %d\n", *reply)
	return nil
}

// Addition is the method that the client will call
func (c *Calculator) Addition(args *Args, reply *int) error {
	if args == nil {
		return errors.New("arguments cannot be nil")
	}
	fmt.Printf("Arguments received: %d + %d\n", args.A, args.B)
	*reply = args.A + args.B
	fmt.Printf("Sending response to client: %d\n", *reply)
	return nil
}

// Subtraction is the mthod that the client will call
func (c *Calculator) Subtraction(args *Args, reply *int) error {
	if args == nil {
		return errors.New("arguments cannot be nil")
	}
	fmt.Printf("Arguments received: %d - %d\n", args.A, args.B)
	*reply = args.A - args.B
	fmt.Printf("Sending response to client: %d\n", *reply)
	return nil
}

func main() {
	calc := new(Calculator)
	rpc.Register(calc) // Register te Calculator service

	listener, err := net.Listen("tcp", ":8080") // Listen on port 8080
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")
	rpc.Accept(listener) // Accept incoming RPC connections
}

package main

import (
	"fmt"
	"net/rpc"
)

// Arguments to pass to the remote method
type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080") // Connect to server at localhost:8080
	if err != nil {
		fmt.Println("Error connecting: ", err)
		return
	}
	defer client.Close()

	fmt.Println("Operation Menu")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Choose one operation: ")

	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Error reading choice: ", err)
		return
	}

	// Prepare arguments
	var args Args
	var result int
	switch choice {
	case 1:
		fmt.Print("Enter first number: ")
		_, err = fmt.Scanln(&args.A)
		if err != nil {
			fmt.Println("Error reading first number: ", err)
			return
		}
		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&args.B)
		if err != nil {
			fmt.Println("Error reading second number: ", err)
			return
		}
		// Call Addition method on the server
		err = client.Call("Calculator.Addition", args, &result)
		if err != nil {
			fmt.Println("Error calling Calculator.Addition: ", err)
			return
		}
		fmt.Printf("Result of Addition: %d + %d = %d\n", args.A, args.B, result)
	case 2:
		fmt.Print("Enter first number: ")
		_, err = fmt.Scanln(&args.A)
		if err != nil {
			fmt.Println("Error reading first number: ", err)
			return
		}
		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&args.B)
		if err != nil {
			fmt.Println("Error reading second number: ", err)
			return
		}
		// Call Addition method on the server
		err = client.Call("Calculator.Subtraction", args, &result)
		if err != nil {
			fmt.Println("Error calling Calculator.Subtraction: ", err)
			return
		}
		fmt.Printf("Result of Subtraction: %d - %d = %d\n", args.A, args.B, result)
	case 3:
		fmt.Print("Enter first number: ")
		_, err = fmt.Scanln(&args.A)
		if err != nil {
			fmt.Println("Error reading first number: ", err)
			return
		}
		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&args.B)
		if err != nil {
			fmt.Println("Error reading second number: ", err)
			return
		}
		// Call Addition method on the server
		err = client.Call("Calculator.Multiply", args, &result)
		if err != nil {
			fmt.Println("Error calling Calculator.Multiply: ", err)
			return
		}
		fmt.Printf("Result of Multiplication: %d * %d = %d\n", args.A, args.B, result)
	case 4:
		fmt.Print("Enter first number: ")
		_, err = fmt.Scanln(&args.A)
		if err != nil {
			fmt.Println("Error reading first number: ", err)
			return
		}
		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&args.B)
		if err != nil {
			fmt.Println("Error reading second number: ", err)
			return
		}
		// Call Addition method on the server
		err = client.Call("Calculator.Divide", args, &result)
		if err != nil {
			fmt.Println("Error calling Calculator.Divide: ", err)
			return
		}
		fmt.Printf("Result of Division: %d / %d = %d\n", args.A, args.B, result)

	default:
		fmt.Println("No operation available for this choice")
	}
}

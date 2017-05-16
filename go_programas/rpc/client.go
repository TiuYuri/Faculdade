package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Q, R int
}

func readArgs() Args {
	var a, b int
	fmt.Println("A: ")
	fmt.Scanln(&a)
	fmt.Println("B: ")
	fmt.Scanln(&b)
	return Args{a, b}
}
func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}

}
func main() {
	service := "localhost:4321"
	client, err := rpc.Dial("tcp", service)
	defer client.Close()
	checkError("Dial: ", err)
	fmt.Println("* - multiplicação")
	fmt.Println("/ - divisão")
	fmt.Println("+ - soma")
	fmt.Println("- - subtração")
	var op byte
	fmt.Scanf("%c\n", &op)
	switch op {
	case '*':
		args := readArgs()
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		checkError("Multiply: ", err)
		fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
		os.Exit(0)
	case '-':
		args := readArgs()
		var reply int
		err = client.Call("Arith.Subtraction", args, &reply)
		checkError("Subtraction: ", err)
		fmt.Printf("%d - %d = %d\n", args.A, args.B, reply)
		os.Exit(0)
	case '+':
		args := readArgs()
		var reply int
		err = client.Call("Arith.Sum", args, &reply)
		checkError("Sum: ", err)
		fmt.Printf("%d + %d = %d\n", args.A, args.B, reply)
		os.Exit(0)
	case '/':
		args := readArgs()
		var reply Quotient
		err = client.Call("Arith.Divide", args, &reply)
		checkError("Divide: ", err)
		fmt.Printf("%d / %d = (%d,%d)\n", args.A, args.B, reply.Q, reply.R)
		os.Exit(0)
	default:
		fmt.Println("Opção inválida: ", op)
		os.Exit(1)
	}
}

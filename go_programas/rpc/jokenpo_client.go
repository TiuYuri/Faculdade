package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Arith int

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
	fmt.Println("Digite: 1 para pedra, 2 para papel ou 3 para tesoura")
	var op int
	fmt.Scanf("%d\n", &op)
	var reply int
	err = client.Call("Arith.Jogar", op, &reply)
	checkError("Jogar: ", err)
	switch {
	case (op - reply) == 0:
		fmt.Println("Empate!")
	case (op-reply) == -2 || (op-reply) == 1:
		fmt.Println("Vitoria!")
	default:
		fmt.Println("Derrota!")
	}
}

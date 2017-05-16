package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Livro struct {
	Id         int
	Titulo     string
	Autor      string
	Isbn       int
	AnoPub     int
	Emprestado bool
	Usuario    string
}

type Emprestimo struct {
	NomeLivro string
	Nome      string
}

type Teste struct {
	ID int
}

// func admin() {
// 	fmt.Printf("menu: \n- 1 isere novo livro \n- 2 solicita remoção do livro \n- 3 consulta por livro\n- 0 sair")
// 	var op int
// 	fmt.Scanf("%d\n", &op)
// 	switch op {
// 	case 1:
// 		a := Livro{}
// 		var var1 int
// 		var var2 string
// 		var reply bool
// 		fmt.Printf("escreva a id do livro: ")
// 		fmt.Scanf("%d\n", &var1)
// 		a.Id = var1
// 		fmt.Printf("escreva o titulo do livro: ")
// 		fmt.Scanf("%s\n", &var2)
// 		a.Titulo = var2
// 		fmt.Printf("escreva o autor do livro: ")
// 		fmt.Scanf("%s\n", &var2)
// 		a.Autor = var2
// 		fmt.Printf("escreva o ano de publicação do livro: ")
// 		fmt.Scanf("%d\n", &var1)
// 		a.AnoPub = var1
// 		fmt.Printf("escreva o isbn do livro: ")
// 		fmt.Scanf("%d\n", &var1)
// 		a.Isbn = var1
// 		a.Emprestado = false
// 		err := Client.Call("Biblioteca.Adicionar", a, &reply)
// 		if err == nil {
// 			if reply == true {
// 				fmt.Println("livro adicionado com sucesso")
// 			} else {
// 				fmt.Println("ops, houve algum erro")
// 			}
// 		} else {
// 			fmt.Println(err)
// 		}
// 	case 2:
// 		var a int
// 		var reply bool
// 		fmt.Printf("escreva a id do livro: ")
// 		fmt.Scanf("%d\n", &a)
// 		err := Client.Call("Biblioteca.Remover", a, &reply)
// 	case 3:
// 		var a string
// 		var reply bool
// 		fmt.Printf("escreva o nome do livro: ")
// 		fmt.Scanf("%s\n", &a)
// 		err := Client.Call("Biblioteca.Busca", a, &reply)
// 	case 0:
// 		return
// 	default:
// 		fmt.Printf("numero não reconhecido")
// 	}
// }

func user() {
	fmt.Printf("menu: \n- 1 solicita um emprestimo \n- 2 solicita devolução \n- 3 consulta por livro\n- 0 sair")
	// var op int
	// fmt.Scanf("%d\n", &op)
	// switch op {
	// case 1:
	// case 2:
	// case 3:
	// case 0:
	// default:
	// }
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}

}
func main() {
	service := "localhost:4321"
	Client, err := rpc.Dial("tcp", service)
	defer Client.Close()
	checkError("Dial: ", err)
	for {
		fmt.Println("Digite: \n- 1 para adiministrador \n- 2 para usuario \n- 0 para sair\n")
		var op int
		fmt.Scanf("%d\n", &op)
		switch {
		case op == 1:
			//admin()
			fmt.Printf("menu: \n- 1 isere novo livro \n- 2 solicita remoção do livro \n- 3 consulta por livro\n- 0 sair\n")
			var op int
			fmt.Scanf("%d\n", &op)
			switch op {
			case 1:
				a := Livro{}
				var var1 int
				var var2 string
				var reply bool
				fmt.Printf("escreva a id do livro: ")
				fmt.Scanf("%d\n", &var1)
				a.Id = var1
				fmt.Printf("escreva o titulo do livro: ")
				fmt.Scanf("%s\n", &var2)
				a.Titulo = var2
				fmt.Printf("escreva o autor do livro: ")
				fmt.Scanf("%s\n", &var2)
				a.Autor = var2
				fmt.Printf("escreva o ano de publicação do livro: ")
				fmt.Scanf("%d\n", &var1)
				a.AnoPub = var1
				fmt.Printf("escreva o isbn do livro: ")
				fmt.Scanf("%d\n", &var1)
				a.Isbn = var1
				a.Emprestado = false
				err := Client.Call("Adicionar", a, &reply)
				if err == nil {
					if reply == true {
						fmt.Println("livro adicionado com sucesso")
					} else {
						fmt.Println("ops, houve algum erro")
					}
				} else {
					fmt.Println(err)
				}
			case 2:
				var a int
				var reply bool
				fmt.Printf("escreva a id do livro: ")
				fmt.Scanf("%d\n", &a)
				err := Client.Call("Biblioteca.Remover", a, &reply)
				checkError("Dial: ", err)
			case 3:
				var a Teste
				var reply Livro
				fmt.Printf("escreva o ID: ")
				fmt.Scanf("%d\n", &a.ID)
				err := Client.Call("Busca", a, &reply)
				checkError("Dial: ", err)
				println(reply.Titulo)
			case 0:
				return
			default:
				fmt.Printf("numero não reconhecido")
			}
		case op == 2:
			user()
		case op == 0:
			return
		default:
			fmt.Println("voce digitou algo errado")
		}
	}
	// var reply int
	// err = client.Call("Arith.Jogar", op, &reply)
	// checkError("Jogar: ", err)
	// switch {
	// case (op - reply) == 0:
	// 	fmt.Println("Empate!")
	// case (op-reply) == -2 || (op-reply) == 1:
	// 	fmt.Println("Vitoria!")
	// default:
	// 	fmt.Println("Derrota!")
	// }
}

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

//** Types **//
type Biblioteca struct {
	Livros map[int]*Livro
}

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

//** END Types **//

//** Setup to library **//
func (B *Biblioteca) Setup() bool {
	B.Livros = make(map[int]*Livro)
	return true
}

func (B *Biblioteca) Inicia() bool {
	var teste bool
	B.Adicionar(&Livro{1, "The Nameless City", "H.P. Lovecraft", 1234, 1996, false, ""}, &teste)
	B.Adicionar(&Livro{2, "teste1", "murilo", 1234, 1996, false, ""}, &teste)
	B.Adicionar(&Livro{3, "teste2", "murilo2", 1234, 1996, false, ""}, &teste)
	B.Adicionar(&Livro{4, "teste3", "murilo3", 1234, 1996, false, ""}, &teste)
	B.Adicionar(&Livro{5, "teste4", "murilo4", 1234, 1996, false, ""}, &teste)
	B.Adicionar(&Livro{6, "teste5", "murilo5", 1234, 1996, false, ""}, &teste)
	return true
}

//** END Setup **//

//** library functions **//

func (B *Biblioteca) Adicionar(a *Livro, reply *bool) error {
	B.Livros[a.Id] = a
	println(a)
	*reply = true
	return nil
}

// func (B *Biblioteca) Remove(a *Emprestimo, reply *bool) error {
// 	_, ok := B.Livros[a.NomeLivro]
// 	if ok {
// 		delete(B.Livros, a.NomeLivro)
// 		*reply = true
// 		return nil
// 	}
// 	*reply = false
// 	return nil
// }

func (B *Biblioteca) Busca(a *Teste, reply *Livro) error {
	i, ok := B.Livros[a.ID]
	if ok {
		reply.Titulo = i.Titulo // ERROR
		return nil
	}
	*reply = Livro{}
	return nil
}

// func (B *Biblioteca) Emprestar(a *Emprestimo, reply *bool) error {
// 	i, ok := B.Livros[a.NomeLivro]
// 	if ok && i.Emprestado == false {
// 		i.Emprestado = true
// 		i.Usuario = a.Nome
// 		B.Livros[a.NomeLivro] = i
// 		*reply = true
// 		return nil
// 	}
// 	*reply = false
// 	return nil
// }

// func (B *Biblioteca) Devolver(a *Emprestimo, reply *bool) error {
// 	i, ok := B.Livros[a.NomeLivro] //mudar o Mapa aqui
// 	if ok && i.Emprestado == true && a.Nome == i.Usuario {
// 		i.Emprestado = false //mudar o Mapa aqui
// 		i.Usuario = "nil"
// 		B.Livros[a.NomeLivro] = i
// 		*reply = true
// 		return nil
// 	}
// 	*reply = false
// 	return nil
// }

//** END Library functions **//

//** DO NOT CHANGE **//
func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}

func main() {
	//** Initialization **//
	B := Biblioteca{}
	B.Setup()
	B.Inicia()

	//** Tests **//
	var testoso Teste
	testoso.ID = 1
	var livrin Livro
	B.Busca(&testoso, &livrin)
	print(livrin.Id)
	print("busco")
	print(livrin.Titulo)
	//** END TESTS **//

	//** network **//
	tcpAddr, err :=
		net.ResolveTCPAddr("tcp", "localhost:4321")
	checkError("ResolveTCPAddr: ", err)
	listener, err :=
		net.ListenTCP("tcp", tcpAddr)
	checkError("ListenTCP: ", err)
	for {
		conn, err := listener.Accept()
		checkError("Accept: ", err)
		go rpc.ServeConn(conn)
	}
	// colocar o mapa na file
}

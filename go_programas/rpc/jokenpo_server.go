package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Arith int

func (t *Arith) Jogar(a *Arith, reply *int) error {
	rand.Seed(time.Now().Unix())
	*reply = rand.Intn(3)
	return nil
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
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
}

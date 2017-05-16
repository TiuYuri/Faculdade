package main

import (
	"io"
	"os"
	"fmt"
	"net"
	"time"
	)
func main () {
	listener, err := net.Listen("tcp", "localhost:13000")
	if err != nil {
		fmt.Println("Erro no Listen: ", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro no Accept: ", err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn (conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String() + "\n"
	io.WriteString(conn, daytime)
}
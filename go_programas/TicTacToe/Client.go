package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
)

type client chan <- string

var (
	entering = make(chan client)	
	leaving = make(chan client)
	messages = make(chan string)
)

func main () {
	listener, err := net.Listen("tcp","localhost:8080")
	if err != nil {
		fmt.Println("Erro no Listen: ", err)
		os.Exit(1)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro no Accept: ", err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster () {
	clients := make(map[client]bool)
	for {
		select {
			case msg := <- messages:
			for cli := range clients {
				cli <- msg
			}
			case cli := <- entering:
			clients[cli] = true
			case cli := <- leaving:
			delete(clients, cli)
					close(cli)
		}
	}
}

func handleConn (conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "ID: " + who
	messages <- who + " entrou"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " saiu"
	conn.Close()
}

func clientWriter (conn net.Conn,
	ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
package main

import (
	"os"
	"fmt"
	"net"
	"time"
)

func main () {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:6666")
	if err != nil {
		fmt.Println("Erro no ResolveUDPAddr: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Erro no ListenUDP: ", err)
		os.Exit(1)
	}
	for {
    	handleConn(conn)
    }
}

func handleConn (conn *net.UDPConn) {
	var buf [1024]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println("Erro no ReadFromUDP: ", err)
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
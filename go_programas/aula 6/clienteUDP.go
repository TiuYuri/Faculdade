package main

import (
	"os"
	"fmt"
	"net"
)

func main () {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:6666")	
	checkError("Erro no ResolveUDPAddr:", err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError("Erro no Dial:", err)
	_, err = conn.Write([]byte("data e hora, por favor"))
	checkError("Erro no Write:", err)
	var reply [1024]byte
	n, err := conn.Read(reply[0:])
	checkError("Erro no Read:", err)
	fmt.Println(string(reply[0:n]))
}
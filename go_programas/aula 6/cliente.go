package main

import "io"
import "os"
import "fmt"
import "net"

func main () {
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		fmt.Println("Erro no Dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
	}
	func mustCopy (dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Erro no mustCopy: ", err)
		os.Exit(1)
	}
}
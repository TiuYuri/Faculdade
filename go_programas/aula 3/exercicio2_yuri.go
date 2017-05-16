package main

import (
	"fmt";
	"os";
	"bufio";
	"strings";
 )

func freq (scanner *bufio.Scanner) map[string]int {// recebe o scaner do bufio
	var m map[string]int = make(map[string]int) //mapa
	for scanner.Scan() {//lendo token por token
		x := strings.Fields(scanner.Text()) //separa a linha
		for i := 0; i < len(x); i++ { // separa as palavra
		valor, ok := m[x[i]] // valor(se existir) ok = true ou false(existe ou n)
			if ok {
				m[x[i]] = valor + 1 
			} else {
				m[x[i]] = 1
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
	return m
}

func main () {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run freq.go <nome-do-arquivo>")
		os.Exit(1)
	}
	filename := os.Args[1]
	input, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	m := freq(scanner)
	for k, v := range(m) {
		fmt.Println(k, v)
	}
	os.Exit(0)
}
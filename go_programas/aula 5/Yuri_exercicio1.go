package main

import  ( "fmt"
		//  "strings"
		)
func conta_vogais (s string, c chan int) {
	soma := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'o' || s[i] == 'i' || s[i] == 'u'{
			soma++
		}
	}
	c <- soma
}

func main () {
	s := "yuri oliveira franco"
	c := make(chan int)
	meio := len(s) / 2
	go conta_vogais(s[:meio], c)
	go conta_vogais(s[meio:], c)
	x, y := <- c, <- c
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
}
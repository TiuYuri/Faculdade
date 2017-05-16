package main

import (
	"fmt";
)

func iter (s []int) func () (int, bool){
	i:= -1
	return func () (int, bool){
		i++
		if i < len(s){
			return s[i], true
		} else {
			return 0, false
		}
	}
}

func main () {
var v = []int { 1, 2, 3, 4, 5 }
	i := iter(v)
	for e, ok := i(); ok == true; e, ok = i() {
		fmt.Println(e)
	}
}
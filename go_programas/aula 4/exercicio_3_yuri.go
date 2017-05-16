package main

import (
 "fmt"
 "log"
 "net/http"
)

type Struct struct {
	Saudacao string
	Fulano string
}

type String string

func (s String) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	for i:=0; i < 1000; i++{
		fmt.Fprint(w, "Olá,	Mundo!")
	}
}

func (t Struct) ServeHTTP (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, t.Saudacao, t.Fulano)
}

func main () {
	var s String
	var t Struct
	t.Saudacao ="olá " 
	t.Fulano = "Fulano"
	
	http.Handle("/string", s)
	http.Handle("/struct", t)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
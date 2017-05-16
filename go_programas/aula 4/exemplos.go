package main

//import "fmt"
/*
type Ponto struct {
	x, y float64
}

func (p Ponto) Ponto2String () string {
	return fmt.Sprintf("{x = %.1f, y = %.1f}", p.x, p.y)
}

func main () {
	p := Ponto{0.0, 5.0}
	fmt.Println(p.Ponto2String())
}
*/

/*type Ponto struct {
	x, y float64
}

func Ponto2String (p Ponto) string {
	return fmt.Sprintf("{x = %.1f, y = %.1f}", p.x, p.y)
}

func main () {
	p := Ponto{0.0, 5.0}
	fmt.Println(Ponto2String(p))//mudança no parametro
}*/

/*type MyString string

func (s MyString) Print () {
	fmt.Printf("MyString = %s\n", s)
}

func main () {
	s := MyString("teste 1 2 3")
	s.Print()
}*/


/*type Ponto struct {
	x, y float64
}

func (p Ponto) Ponto2String () string {
	return fmt.Sprintf("{x = %.1f, y = %.1f}", p.x, p.y)
}

func (p *Ponto) Move (x, y float64) {
	p.x += x
	p.y += y
}

func main () {
	p := Ponto{0.0, 5.0}
	p.Move(10.0, 5.0)
	fmt.Println(p.Ponto2String())
}*/

/*type Ponto struct {
	x, y float64
}

func (p Ponto) String () string {
	return fmt.Sprintf("{x = %.1f, y = %.1f}", p.x, p.y)
}

func (p *Ponto) Move (x, y float64) {
	p.x += x
	p.y += y
}

func main () {
	p := Ponto{0.0, 5.0}
	fmt.Println("Antes de mover p =", p)
	p.Move(10.0, 5.0)
	fmt.Println("Depois de mover p =", p)
}*/

/*type ErroDivisaoPorZero int

func (e ErroDivisaoPorZero) Error () string {
	return "divisão por zero"
}

func idiv (x int, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, ErroDivisaoPorZero(x)
	} else {
		return x / y, x % y, nil
	}
}

func main () {
	x, y := 9, 2
	q, r, err := idiv(x,y)
	if err == nil {
		fmt.Printf("%d/%d = (%d,%d)\n", x, y, q, r)
	} else {
		fmt.Println("Erro:", err)
	}
}*/

/*import "io"
import "os"
import "strings"

type rot13Reader struct {
	r io.Reader
}

func (self *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = self.r.Read(p)
	for i, b := range(p) {
		if b <= 'Z' && b >='A' {
			p[i] = (b - 'A' + 13) % 26 + 'A'
		} else if b >= 'a' && b <= 'z' {
			p[i] = (b - 'a' + 13) % 26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Qrpvsenfgr b frterqb!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}*/

import (
"fmt"
"log"
"net/http"
)

type String string

func (s String) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "VAI SE FODER\n")
}

func main () {
	var s String
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal(err)
	}
}
package main
import "fmt"

type Ponto struct {
    x, y float64
}

func (p *Ponto) Ponto2String () string {
    return fmt.Sprintf("{x = %.1f, y = %.1f}", p.x, p.y)
}
func (p Ponto) Move (x, y float64) {
    p.x += x
    p.y += y
}

func main () {
    p := Ponto{0.0, 5.0}
    p.Move(10.0, 5.0)
    fmt.Println(p.Ponto2String())
}
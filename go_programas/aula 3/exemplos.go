package main

import "fmt"

/*func main () {
	i, j := 1, 2
	var p1 *int = &i
	p2 := &j
	*p1, *p2 = *p2, *p1 // = *p1, *p2 = 2, 1 avalia todo o lado direito antes
	fmt.Println(i, j)
}*/
/*type Shape struct {
	x int
	y int
}

func main () {
	s := Shape{10, 20}
	fmt.Println(s, s.x, s.y)
}*/

/*type Shape struct {
	x int
	y int
}

func main () {
	s := Shape{10, 20}
	fmt.Println(s, s.x, s.y)
	var p *Shape = &s
	p.x = 5
	p.y = 12
	fmt.Println(p, p.x, p.y)
}*/

//vetor estatico

/*func main () {
	var v [5]int
	for i := 0; i < 5; i++ {
		v[i] = i + 1
		fmt.Printf("v[%d] = %d\n", i, v[i])
	}
}*/

// slice (vetor dinamico)

/*func main () {
	p := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] = %d\n", i, p[i])
	}
}*/

//matriz

/*func main () {
	matriz := [][]int {
		[]int{0},
		[]int{1,2},
		[]int{3,4,5},
	}
	fmt.Println(matriz)
}*/

/*func main () {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("p[2:7] = ", p[2:7]) //pega o array da posição 2 a 6
	fmt.Println("p[:5] ~ p[0:5] =", p[:5]) // do 0 ao 4
	fmt.Println("p[4:] ~ p[4:len(p)] =", p[4:]) // do 4 ao final da lista
}*/

/*func printSlice (name string, slice []int) {
	fmt.Printf("len(%s) = %d, cap(%s) = %d, %v\n",
	name, len(slice), name, cap(slice), slice)
}

func main () {
	v1 := make([]int, 5)
	printSlice("v1", v1)
	v2 := make([]int, 0, 5) // tem 5 posições "alocadas" mas nada inicializado
	printSlice("v2", v2)
	v3 := v1[:3]
	printSlice("v3", v3)
	v4 := v3[3:] 
	printSlice("v4", v4)
}*/

/*func main () {
	var v []int
	if v == nil {
		fmt.Println("nil", len(v), cap(v))
	}
}*/

/*func printSlice (name string, slice []int) {
	fmt.Printf("len(%s) = %d, cap(%s) = %d, %v\n",
	name, len(slice), name, cap(slice), slice)
}

func main () {
	var v []int
	printSlice("v", v)
	v = append(v, 1)
	printSlice("v", v)
	v = append(v, 2)
	printSlice("v", v)
	v = append(v, 3)
	printSlice("v", v)
	v = append(v, 4, 5, 6)
	printSlice("v", v)
	v = append(v, 7, 8, 9)
	printSlice("v", v)
}*/

/*func main () {
	var vogais = []int{'a', 'e', 'i', 'o', 'u'}
	for i, v := range vogais {
		fmt.Printf("vogais[%d] = %d ~ %c\n", i, v, v) // %d tabela ascii
	}
}*/

/*func main () {
	var vogais = []int{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vogais { // _ para não iniciar uma variavel
		fmt.Printf("%c = %d\n", v, v)
	}
}*/

/*import "os"

func main () { // le a linha de comando
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}*/

/*import "strconv"

func main () {
	x := strconv.Itoa(123)
	y, err := strconv.Atoi("456") // so converte coisas que parecem numeros inteiros, senão retorna erro
	fmt.Println(x, y, err)
}*/

/*type Coordenadas struct {
	Latitude, Longitude float64
}

func main () {
	var m map[string]Coordenadas =
	make(map[string]Coordenadas)
	m["PUCPR"] = Coordenadas{ -25.4516088,
							  -49.2527509 } // insere 
	fmt.Println(m)
	delete(m, "PUCPR")
	valor, ok := m["PUCPR"] // valor(se existir) ok = true ou false(existe ou n)
	fmt.Println(valor, ok)
}*/


/*func main () {
	var suc func (x int) int
	soma := func (x, y int) int {
		return x + y
	}
	suc = func (x int) int {
	return x + 1
	}
	fmt.Println(soma(suc(1), suc(1)))
}*/

func newCounter () func () int {
	i := 0
	return func () int {
		i += 1
		return i
	}
}

func main () {
	c1, c2 := newCounter(), newCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c1())
	fmt.Println(c2())
}
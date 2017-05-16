/*package main

import "fmt"

var MAX int = 5

func f (n int) {
	for i := 0; i < MAX; i++ {
		fmt.Println(n, ":", i)
	}
}
func main () {
	go f(0)
	var input string
	fmt.Scanln(&input)
} */

/*package main

import (
	"fmt"
	"time"
	"math/rand"
)
	var MAX int = 5

func f (n int) {
	for i := 0; i < MAX; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main () {
	for i := 0; i < MAX; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}*/

/*package main

import "fmt"

func soma (v []int, c chan int) {
	soma := 0
	for i := 0; i < len(v); i++ {
		soma += v[i]
	}
	c <- soma
}

func main () {
	v := []int{ 1, 2, 3, 4, 5, 6 }
	c := make(chan int)
	meio := len(v) / 2
	go soma(v[:meio], c)
	go soma(v[meio:], c)
	x, y := <- c, <- c
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
}*/

/*package main

import "fmt"

func main () {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)
}*/

package main

import (
	"fmt"
	"sync"
	"time"
)
const N = 100000000

var (
	cont int = 0
	mu sync.Mutex
)

func Inc () {
	mu.Lock()
	cont++
	mu.Unlock()
}

func Valor () int {
	mu.Lock()
	defer mu.Unlock()
	return cont
}

func main () {
	go func () {
		for i := 0; i < N; i++ {
			Inc()
		}
	}()
	for i := 0; i < N; i++ {
		Inc()
	}
	time.Sleep(time.Second)
	fmt.Println(Valor())
}
package main

import 
	("fmt"
	"sync"
	)
const N = 1000000

var saldo int

var (
	cont int = 0
	mu sync.Mutex
)

func Deposito (v int) {
		mu.Lock()
		defer mu.Unlock()
		saldo = saldo + v
	}

	func Saque (v int) {
		mu.Lock()
		defer mu.Unlock()
		saldo = saldo - v
	}
	func Saldo () int {
		mu.Lock()
		defer mu.Unlock()
		return saldo
	}

func main () {
	fim := make(chan bool)
	go func (fim chan <- bool) {
		for i := 1; i <= N; i++ {
			Saque(N * i)
		}
		fim <- true
	}(fim)
	for i := 1; i <= N; i++ {
		Deposito(N * i)
	}
	_ = <- fim
	fmt.Println(Saldo())
}
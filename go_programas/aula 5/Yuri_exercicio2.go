package main

import 
   ("fmt"
   "os"
   "strconv"
	)


func main() {

	if len(os.Args) != 2{
		fmt.Printf("coloque apenas 1 numero")
		return
	}

	//n,_ := strconv.Atoi(os.Args[1])
	

	s,_ := strconv.Atoi(os.Args[1])
	c := make(chan int, s)
	go primos(cap(c), c)
	for i:= range c{
		fmt.Printf("%d | ", i)
	}
}

func primos (n int, c chan int){
    b := make([]bool, n)
    for i := 2; i < n; i++ {
        if b[i] == true { continue }
        c <- i
        for k := i * i; k < n; k += i {
            b[k] = true
        }
    }
	close(c)
}
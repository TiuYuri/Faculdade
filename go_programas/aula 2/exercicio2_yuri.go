package main

import (
	"fmt" 
	"math/rand"
	"time"
)
func main(){
	pedra_papel_tesoura()
}


func pedra_papel_tesoura(){
	var input int
	fmt.Print("entre com o numero: ")
    fmt.Scanln(&input)
	rand.Seed(time.Now().Unix())
	var pc = rand.Intn(3)
	if(pc == 0){
		fmt.Println("pc joga Pedra")
	} else{
		if(pc == 1){
			fmt.Println("pc joga Papel")
		}else{
			fmt.Println("pc joga Tesoura")
		}
	}
	 switch {
		case (input - pc) == 0:
			 fmt.Println("Empate!")
		case (input - pc) == -2 || (input - pc) == 1:
			fmt.Println("Vitoria!")
		default:
			fmt.Println("Derrota!")
	 }
}
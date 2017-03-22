package main

import (
	"fmt" 
	"math"
	)

func main () {

	fmt.Println(calcula_pi(50000))
	fmt.Println(math.Pi)
}

func calcula_pi(precisao int) float64{
	var pi = 0.0
	for i:= 0; i <= precisao; i++{
		pi = pi + math.Pow(-1.0, float64(i))/(2.0*float64(i) + 1.0)
	}
	return (pi*4)
}
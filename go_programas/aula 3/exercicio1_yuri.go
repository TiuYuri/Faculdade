package main
import (
	"fmt";
	"os";
	"strconv";
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("poucos argumentos, saindo...")
		return 
	}
	if len(os.Args) > 3 {
		fmt.Println("muitos argumentos (só avisando)")
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	//fmt.Println("erros: ", err1, err2)
	if err1 != nil{
		fmt.Println("primeiro argumento invalido")
	} else {
		if err2 != nil{
			fmt.Println("segundo argumento invalido")	
		} else {
			fmt.Printf("x = %d, Y= %d \n x + y = %d \n", x, y, (x+y))
		}
	}
}
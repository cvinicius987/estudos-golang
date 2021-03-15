package main

import "fmt"

func main() {

	x := 10

	//Condição boolean
	if x < 100 {
		fmt.Println("Menor que 100")
	}

	//Operador de negação
	if !(x > 100) {
		fmt.Println("Menor que 100")
	}

	//Com declaração, else if e else
	if y := 9; y == 10 {
		fmt.Println("y é IGUAL a 10")
	} else if y > 10 {
		fmt.Println("y MAIOR que 10")

	} else {
		fmt.Println("y MENOR que 10")
	}
}

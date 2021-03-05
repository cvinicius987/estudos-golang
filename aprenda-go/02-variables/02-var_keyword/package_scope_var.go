package main

import (
	"fmt"
)

//Variavel declarada em escopo de package
var initial = 100

func plus(value int) {
	fmt.Printf("Valor da Soma é: %d \n", initial+value)
}

func main() {


	fmt.Printf("Valor Inicial: %d \n", initial)

	var x = 10

	plus(x)
}

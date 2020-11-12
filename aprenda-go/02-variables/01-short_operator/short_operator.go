package main

import (
	"fmt"
)

func main() {

	//declarar variaveis onde o tipo será definido pelo compilador
	//so pode ser utilizado em escopo de função
	//So pode ser utilizado em declaração de uma nova variavel
	x := 10
	name := "Golang"

	fmt.Printf("x: %d, type: %T \n", x, x)
	fmt.Printf("name: %s, type: %T \n", name, name)
}

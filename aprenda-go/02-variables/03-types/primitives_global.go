package main

import (
	"fmt"
)

var name string

func main() {

	//Não é valido pois a declaração especifica uma string
	//name = 10
	name = "Golang é uma linguagem fortemente tipada"

	fmt.Println(name)
}

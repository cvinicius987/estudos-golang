package main

import (
	"fmt"
)

var name string

func main() {

	//Não é valido setar um valor de outro tipo, pois a declaração especifica uma string
	name = "Golang é uma linguagem fortemente tipada"

	fmt.Println(name)
}

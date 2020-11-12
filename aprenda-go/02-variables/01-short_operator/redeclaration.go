package main

import "fmt"

func main() {

	x := "Nome"

	fmt.Println(x)

	//O operador := não pode ser usado para atribuir valores para uma variavel existente
	//x := "Golang"

	x = "Novo nome"

	fmt.Println(x)
}

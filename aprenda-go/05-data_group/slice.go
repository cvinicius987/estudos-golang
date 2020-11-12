package main

import "fmt"

/*
	Slice é um tipo de dado composto

	Um slice internamente é um array;

	Ao declarar um slice não precisamos definir seu tamanho em elementos;
*/
func main() {

	//Declarou um Slice
	languages := []string{"Golang", "Java"}

	fmt.Println(languages)

	//Add novos itens no slice, se fosse array não seria possivel redimensionar
	languages = append(languages, "Kotlin", "JavaScript")

	fmt.Println(languages)
}

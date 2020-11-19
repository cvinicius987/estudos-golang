package main

import "fmt"

/*
 O Range atravessa o Slice para retornar os elementos
*/
func main() {

	languages := []string{"Java", "Kotlin", "JavaScript", "Golang"}

	for index, value := range languages {
		fmt.Printf("Ranking: %d, Linguagem: %s \n", index, value)
	}

	//Redimensionou o Slice para permitir adicionar um novo elemento
	languages = append(languages, "Elixir")

	fmt.Println("\n")

	//Ao usar _ estamos desprezando o valor da declaração desestruturante
	for _, value := range languages {
		fmt.Printf("Linguagem: %s \n", value)
	}
}

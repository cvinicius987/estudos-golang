package main

import "fmt"

func main() {

	languages := []string{"Java", "Kotlin", "JavaScript", "Golang", "Elixir"}

	//Recupera itens do slices onde analisa os itens do elemento
	//o ultimo item é desprezado
	java_kotlin := languages[0:2]

	fmt.Println("Java e Kotlin: ", java_kotlin, "\n")

	//Recupera todo o conteudo do slice
	all_languages := languages[:]

	fmt.Println(all_languages, "\n")

	//Recupera todos os itens até o elemento da direita
	java_kotlin_js := languages[:3]

	fmt.Println(java_kotlin_js, "\n")

	//Iterando sem Range e recuperando de um Slice
	for index := 0; index < len(languages); index++ {

		fmt.Printf("Index: %d - Linguagem: %v\n", index, languages[index])
	}
}

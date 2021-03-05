package main

import "fmt"

func main() {

	languages := []string{"Java", "Kotlin", "JavaScript", "Golang", "Elixir"}

	//Iterando sem Range e recuperando de um Slice
	for index := 0; index < len(languages); index++ {

		fmt.Printf("Index: %d - Linguagem: %v\n", index, languages[index])
	}

	fmt.Print("\n =============== \n")

	//Iterando com rage
	for index, _ := range languages {

		fmt.Printf("Index: %d - Linguagem: %v\n", index, languages[index])
	}

}

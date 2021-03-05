package main

import "fmt"

func main() {

	list := []string{"Java", "Kotlin", "Go"}

	//Função anonima é utilizada somente uma vez
	func(items []string) {

		for index, value := range items {
			fmt.Println(index, " = ", value)
		}

	}(list)

	//Função anonima que pode ser usado mais de uma vez
	functionTeste := func(items []string) {

		for index, value := range items {
			fmt.Println(index, " = ", value)
		}

	}

	functionTeste(list)
}

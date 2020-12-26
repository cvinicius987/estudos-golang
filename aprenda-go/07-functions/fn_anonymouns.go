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
}

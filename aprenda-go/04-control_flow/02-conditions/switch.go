package main

import "fmt"

func main() {

	//Basic
	x := 10

	switch {
	case x < 5:
		fmt.Println("x MENOR que 5")
	case x == 5:
		fmt.Println("x IGUAL a 5")
	case x > 5:
		fmt.Println("x MAIOR que 5")
	}

	//Expression
	language := "Golang"

	switch language {
	case "Golang":
		fmt.Println("A linguagem utilizada é Golang")
	default:
		fmt.Println("A linguagem é ", language)
	}
}

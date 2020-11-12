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

	//fallthrough ==>> Executa a expressão abaixo da encontrada
	//é possivel add mais de uma expressão para o case
	car := "uno"

	switch car {
	case "uno":
		fmt.Println("uno")
		fallthrough
	case "gol":
		fmt.Println("gol")
	case "marea":
		fmt.Println("marea")
	case "ka", "fiesta":
		fmt.Println("Carro é um Ford")
	default:
		fmt.Println("Não existe carro")
	}
}

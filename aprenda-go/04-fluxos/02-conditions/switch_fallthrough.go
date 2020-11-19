package main

import "fmt"

func main() {

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

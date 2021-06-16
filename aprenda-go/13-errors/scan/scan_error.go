package main

import "fmt"

func main() {

	var name, language string

	fmt.Print("Digite seu nome: ")

	//Pode receber um erro
	_, err := fmt.Scan(&name)

	//Verifica se existe um erro na chamada
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite sua linguagem: ")

	//Pode receber um erro
	_, err = fmt.Scan(&language)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nO usuario: %s, programa na linguagem: %s \n", name, language)
}

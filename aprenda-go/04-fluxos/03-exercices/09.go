package main

import (
	"fmt"
)

/*
- Crie um programa que utilize a declaração switch, onde o switch statement
seja uma variável do tipo string com identificador "esporteFavorito".
*/
func main() {

	esporteFavorito := "volei"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Esporte é futebol")
	case "natação":
		fmt.Println("Esporte é natação")
	case "basquete":
		fmt.Println("Esporte é basquete")
	default:
		fmt.Println("Não tem esporte favorito")
	}
}

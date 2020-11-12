package main

import "fmt"

/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu
*/
func main() {

	nascimento := 1987
	anoContar := 2100

	for nascimento <= anoContar {

		fmt.Println(nascimento)
		nascimento++
	}
}

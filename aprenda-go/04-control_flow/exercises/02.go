package main

import "fmt"

/*
- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
*/
func main() {

	for i := 65; i <= 90; i++ {

		fmt.Println("\n", i)

		for x := 0; x < 3; x++ {

			fmt.Printf("\t%#U \n", i)
		}

	}

}

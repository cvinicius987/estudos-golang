package main

import "fmt"

/*
- Utilizando a solução anterior, adicione as opções else if e else.
*/
func main() {

	x := 11

	if x > 10 {
		fmt.Println("X é maior que 10")

	} else if x < 10 {
		fmt.Println("X é menor que 10")

	} else {
		fmt.Println(" X é igual a 10.")
	}
}

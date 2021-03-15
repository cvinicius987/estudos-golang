package main

import "fmt"

/*
 Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/
func main() {

	salario := 99

	switch {
	case salario < 100:
		fmt.Println("O salario", salario, "é MENOR que 100")
	case salario == 100:
		fmt.Println("O salario é 100")
	default:
		fmt.Println("O salario", salario, "é MAIOR que 100")
	}
}

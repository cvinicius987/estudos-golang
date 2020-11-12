package main

import (
	"fmt"
)

/*
 * Tipagem estatica, após declarado não é possivel alterar o valor
 */
func main() {

	//Integer
	var number int = 10

	//Float
	var money float64 = 10.55

	//String
	var name string = "Golang"

	//Boolean
	var isOk bool = true

	fmt.Println(number, money, name, isOk)
}

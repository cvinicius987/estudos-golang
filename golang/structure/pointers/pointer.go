package main

import "fmt"

func main() {

	//variavel
	x := "Golang"

	//endereço de x
	pointer_x := &x

	fmt.Println("Var x: ", x, "Endereço de x:", pointer_x)

	fmt.Println("Valor do endereço:", *pointer_x)
}

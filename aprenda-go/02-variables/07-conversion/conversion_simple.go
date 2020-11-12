package main

import "fmt"

type client string

func main() {

	name := "Nome"

	var clientNew client = "Nome do cliente"

	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", clientNew, clientNew)

	//essa conversão não é valida, mesmo o clientNew ser derivado de string
	//é necessario a chamada explicita da conversão
	//name = clientNew

	name = string(clientNew)

	fmt.Printf("%v, %T\n", name, name)
}

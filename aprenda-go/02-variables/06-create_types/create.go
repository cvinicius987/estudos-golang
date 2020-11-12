package main

import "fmt"

//Um tipo é devivado de um outro tipo, mas não pertence a hierarquia
type client string

func main() {

	name := "Nome"

	var clientNew client = "Nome do cliente"

	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", clientNew, clientNew)
}

package main

import "fmt"

/**
Funções com receivers

func (receiver) identifier(parameters) (returns) {
	code
}
*/

type client struct {
	name    string
	address string
}

func (c client) message() {
	fmt.Println("Bem-vindo:", c.name, "você mora em:", c.address)
}

func main() {

	client1 := client{
		name:    "Caio",
		address: "São Paulo",
	}

	client1.message()
}

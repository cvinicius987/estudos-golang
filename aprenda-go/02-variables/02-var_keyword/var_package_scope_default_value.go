package main

import (
	"fmt"
)

//Declarado as variaveis sem inicializar, onde os valores default serão assumidos
var name string
var total int

var price float64

func main() {

	//Inicializa as variaveis
	name  = "Golang"
	total = 10

	fmt.Printf("Name: %s, Total: %d, Price: %v \n", name, total, price)
}

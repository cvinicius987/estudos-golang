package main

import (
	"fmt"
)

func main() {

	//Realiza a declaração de uma variavel definindo seu tipo
	var x = 10

	//Ao definir um tipo, a varaivel é inicializada com valor padrão
	var y string

	fmt.Printf("x: %d, type: %T \n", x, x)
	fmt.Printf("y: %s, type: %T \n", y, y)
}

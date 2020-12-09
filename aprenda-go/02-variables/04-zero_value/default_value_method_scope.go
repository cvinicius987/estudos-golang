package main

import (
	"fmt"
)

func main() {

	//Default: 0
	var number int

	//Erro de compilação, o valor default só será atribuido em escopo local
	//quando o tipo for explicito
	//var number2

	fmt.Printf("%v, %T\n", number, number)
}

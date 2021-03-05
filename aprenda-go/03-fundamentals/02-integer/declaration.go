package main

import "fmt"

func main() {

	//Os tipos inteiros são definidos entre int32 ou int64 pela arquitetura da maquina
	//onde o programa esta sendo executado
	//Por padrão use int e deixe a maquina definir qual é mais performatico
	var x int = 10

	y := 10.0

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
}

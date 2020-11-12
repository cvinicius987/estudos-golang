package main

import "fmt"

/*
Arrays são estruturas de dados para armazenamento de um tipo especifico;

Deve ser declarada setando o numero de posições que serão necessarias;

No dia a dia não é muito utilizada, mas em casos especificos de performance
podemos utilizar no lugar das slices;
*/

func main() {

	var names [4]string

	names[0] = "Golang"
	names[1] = "Java"
	names[2] = "Kotlin"

	fmt.Println(names[0], names[1], names[2], names[3])

	x := [2]int{1, 2}

	fmt.Println(x[0], x[1])
}

package main

import "fmt"

/*
	Todos os valores ficam armazenados na memória.
	- Toda localização na memória possui um endereço.
	- Um pointeiro se refere a esse endereço.
*/

func main() {

	//& => é usado para exibir o endereço da memoria
	//* => é usado para recuperar o value do endereço

	x := "Golang"

	fmt.Println("==== Main")
	fmt.Println("x address:", &x)
	fmt.Println("x value:", x)

	viewMemory(x)

	viewMemoryByPointer(&x)
}

func viewMemory(x string) {

	fmt.Println("==== viewMemory")
	fmt.Println("x address:", &x)
	fmt.Println("x value:", x)
}

func viewMemoryByPointer(x *string) {

	fmt.Println("==== viewMemoryByPointer")
	fmt.Println("x address:", x)
	fmt.Println("x value:", *x)
}

package main

import "fmt"

func main() {

	/*
		String são imutaveis, ou seja, a acada alteração uma nova string será criada
	*/

	s := "Hello String...."

	fmt.Printf("%v, %T", s, s)

	//Converte para bytes
	sb := []byte(s)

	fmt.Printf("%v, %T", sb, sb)
}

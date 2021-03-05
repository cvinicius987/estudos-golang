package main

import "fmt"

func main() {

	s := "Hello String...."

	//Converte para bytes
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("Binario: %b, Type: %T, Valor: %v, Unicode: %#U, Base16: %#x\n", v, v, v, v, v)
	}

	fmt.Printf("%v, %T", sb, sb)
}

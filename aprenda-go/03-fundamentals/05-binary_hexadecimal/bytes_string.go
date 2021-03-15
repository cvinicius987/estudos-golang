package main

import "fmt"

func main() {

	s := "Hello String...."

	//Converte para bytes
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("Unicode: %#U, Decimal (ASC): %v, Binario: %b, Hexadecimal: %#x \n", v, v, v, v)
	}
}

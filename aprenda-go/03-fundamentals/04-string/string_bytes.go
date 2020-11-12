package main

import "fmt"

func main() {

	s := "Hello String...."

	//Converte para bytes
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("%b, %v, %T, %#U, %#x\n", v, v, v, v)
	}

	fmt.Printf("%v, %T", sb, sb)
}

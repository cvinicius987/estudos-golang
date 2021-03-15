package main

import "fmt"

func main() {

	s := `{"Name":"Golang"}`

	fmt.Printf("%v, %T", s, s)

	//Converte para bytes
	sb := []byte(s)

	fmt.Printf("%v, %T", sb, sb)
}

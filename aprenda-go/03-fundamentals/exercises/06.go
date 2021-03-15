package main

import "fmt"

const (
	_ = iota + 2020
	first
	second
	third
	fourth
)

/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/
func main() {

	fmt.Println(first, second, third, fourth)
}

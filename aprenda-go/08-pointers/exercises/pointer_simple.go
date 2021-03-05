package main

import "fmt"

/*
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
*/
func main() {

	x := 10
	y := 10

	language := "Golang"
	language2 := "Golang"

	fmt.Println("x:", &x, "\ny:", &y, "\nlanguage:", &language, "\nlanguage2:", &language2)
}

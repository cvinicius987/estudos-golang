package main

import "fmt"

const country string = "Brasil"

const (
	SP = "São Paulo"
	RJ = "Rio de Janeiro"
)

/*
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
*/
func main() {

	fmt.Println(country)
	fmt.Println(SP, RJ)
}

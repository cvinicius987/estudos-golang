package main

import "fmt"

func main() {

	pessoa := struct {
		name  string
		idade int
	}{
		name:  "Caio",
		idade: 33}

	fmt.Println(pessoa)
	fmt.Println(pessoa.name)
}

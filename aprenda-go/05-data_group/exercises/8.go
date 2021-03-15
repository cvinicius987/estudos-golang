package main

import "fmt"

/*
 Crie um map com key tipo string e value tipo []string.
- Key deve conter nomes no formato sobrenome_nome
- Value deve conter os hobbies favoritos da pessoa
*/
func main() {

	pessoas := map[string][]string{
		"caio_vinicius":    {"natação", "futebol", "videos"},
		"fernanda_machado": {"volei", "filmes", "teatro"},
	}

	for key, value := range pessoas {

		fmt.Println(key)

		for index, hobbie := range value {
			fmt.Println("\t", index, "-", hobbie)
		}
	}
}

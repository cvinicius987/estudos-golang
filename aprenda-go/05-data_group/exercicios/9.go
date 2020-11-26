package main

import "fmt"

/*
- Utilizando o exercício anterior,
adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/
func main() {

	pessoas := map[string][]string{
		"caio_vinicius":    {"natação", "futebol", "videos"},
		"fernanda_machado": {"volei", "filmes", "teatro"},
	}

	//add no map
	pessoas["gabriela_antunes"] = []string{"ensinar", "ler", "correr"}

	for key, value := range pessoas {

		fmt.Println(key)

		for index, hobbie := range value {
			fmt.Println("\t", index, "-", hobbie)
		}
	}
}

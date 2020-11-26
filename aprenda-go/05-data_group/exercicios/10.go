package main

import "fmt"

/*
- Utilizando o exercício anterior,
- Remova um key;
- Exiba o resultado;
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

	//exclusão
	delete(pessoas, "fernanda_machado")

	fmt.Println("\n")

	for key, value := range pessoas {

		fmt.Println(key)

		for index, hobbie := range value {
			fmt.Println("\t", index, "-", hobbie)
		}
	}
}

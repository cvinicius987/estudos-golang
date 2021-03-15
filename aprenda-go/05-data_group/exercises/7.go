package main

import "fmt"

/*
- Crie uma slice contendo slices de strings ([][]string).
Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
func main() {

	conteudo := [][]string{
		[]string{"caio", "vinicius", "natação"},
		[]string{"joão", "augusto", "futebol"},
		[]string{"fernanda", "machado", "volei"},
	}

	for x := 0; x < len(conteudo); x++ {

		fmt.Println("Pessoa", x)

		for y := 0; y < len(conteudo[x]); y++ {
			fmt.Println("\t", conteudo[x][y])
		}
	}
}

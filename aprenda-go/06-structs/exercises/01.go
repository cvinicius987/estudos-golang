package main

import "fmt"

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores,
utilizando range na slice que contem os sabores de sorvete.
*/

type pessoa struct {
	nome      string
	sobrenome string
	sorvetes  []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "caio",
		sobrenome: "vinicius",
		sorvetes:  []string{"chocolate", "flocos"},
	}

	pessoa2 := pessoa{
		nome:      "renata",
		sobrenome: "vasconcellos",
		sorvetes:  []string{"morango", "abacaxi"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, sorvete := range pessoa1.sorvetes {
		fmt.Printf("\t %v\n", sorvete)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, sorvete := range pessoa2.sorvetes {
		fmt.Printf("\t %v\n", sorvete)
	}
}

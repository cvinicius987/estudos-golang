package main

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
  utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range,
 dentro do range anterior.
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

	mapPessoa := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for key, p := range mapPessoa {
		fmt.Println(key)
		fmt.Println("\t", p.nome)

		for _, sorvete := range p.sorvetes {
			fmt.Printf("\t\t %v\n", sorvete)
		}
	}
}

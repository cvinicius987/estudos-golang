package main

import "fmt"

/*
- Crie uma slice usando make que possa conter todos os estados do Brasil.
	- Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará",
	"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
	"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
	"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
	"Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*
*/

var conteudoEstados = [26]string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará",
	"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
	"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
	"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
	"Rondônia", "Roraima", "Santa Catarina", "São Paulo",
	"Sergipe", "Tocantins"}

func main() {

	estados := make([]string, 26, 30)

	for index, c := range conteudoEstados {
		estados[index] = c
	}

	fmt.Println("Len:", len(estados), cap(estados))

	for x := 0; x < len(estados); x++ {

		fmt.Println("Estado:", estados[x])
	}
}

package main

import "fmt"

type pessoa struct {
	name  string
	idade int
}

type professional struct {
	title string
	pessoa
}

func main() {

	pessoa1 := pessoa{"Caio", 33}

	pessoa2 := pessoa{"Claudia", 40}

	professional1 := professional{
		title:  "Programador",
		pessoa: pessoa1,
	}

	professional2 := professional{"Gerente", pessoa2}

	//Acessando por niveis de objetos
	fmt.Println(professional1.title, professional1.pessoa.name, professional1.pessoa.idade)

	//Acessando diretamente de quando os structs não possui nomes correlacionados
	fmt.Println(professional2.title, professional2.name, professional2.idade)
}

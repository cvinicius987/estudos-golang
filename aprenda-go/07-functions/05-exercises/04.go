package main

import (
	"fmt"
)

/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

func (p pessoa) nomeCompleto(){
	fmt.Println("Nome completo:", p.nome, p.sobrenome, "Idade: ", p.idade)
}

func main(){

	p1 := pessoa{
		nome: "Caio",
		sobrenome: "Vinicius",
		idade: 33,
	}

	p1.nomeCompleto()
}
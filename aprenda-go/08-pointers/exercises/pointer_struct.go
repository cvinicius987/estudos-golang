package main

import "fmt"

/*
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors​
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
*/

type pessoa struct {
	name  string
	owner string
}

func main() {

	p1 := pessoa{
		name:  "Go",
		owner: "Google",
	}

	fmt.Println("Main antes de mudar:", p1)

	mudeme(&p1)

	fmt.Println("Main apos mudar:", p1)
}

func mudeme(p *pessoa) {

	(*p).name = "Java"
	p.owner = "Oracle"

	fmt.Println("Mudeme:", *p)
}

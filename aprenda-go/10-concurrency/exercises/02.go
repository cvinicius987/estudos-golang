package main

import (
	"fmt"
	"math/rand"
)

/*
Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/

type Humanos interface {
	falar()
}

type Pessoa struct{}

func (p *Pessoa) falar() {
	fmt.Println("Humano falando um numero", rand.Int())
}

func dizerAlgumaCoisa(humano Humanos) {
	humano.falar()
}

func main() {

	pessoa := Pessoa{}

	dizerAlgumaCoisa(&pessoa)
}

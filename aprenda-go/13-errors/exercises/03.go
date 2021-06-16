package main

import (
	"fmt"
	"log"
)

type erroEspecial struct {
	Message string
}

func (n erroEspecial) Error() string {
	return fmt.Sprintf("Aqui erro no struct: %v", n.Message)
}

/*
- Crie um struct "erroEspecial" que implemente a interface builtin.error.
- Crie uma função que tenha um valor do tipo error como parâmetro.
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/
func main() {

	especial := erroEspecial{"Mensagem de erro"}

	testError(especial)
}

func testError(err error) {

	log.Fatalln(err)
}

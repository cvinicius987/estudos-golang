package main

import "fmt"

/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
func main() {

	materia := struct {
		professores []string
		linguagens  map[string]int
	}{
		professores: []string{"Carlos", "Fernanda"},
		linguagens:  map[string]int{"java": 1, "golang": 2},
	}

	fmt.Println(materia)
}

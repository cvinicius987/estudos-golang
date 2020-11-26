package main

import "fmt"

/*
     - Crie um array que suporte 5 valores to tipo int
	 - Atribua valores aos seus índices
	 - Utilizar range
	 - Monstre o tipo do array
*/
func main() {

	var names [5]string

	names[0] = "java"
	names[1] = "kotlin"
	names[2] = "golang"
	names[3] = "python"
	names[4] = "elixir"

	for index, conteudo := range names {
		fmt.Println("Index:", index, "Conteudo:", conteudo)
	}

	fmt.Printf("%T", names)
}

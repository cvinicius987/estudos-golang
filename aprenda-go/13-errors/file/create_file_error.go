package main

import (
	"fmt"
	"os"
)

const FILE = "aprenda-go/13-errors/file/languages.txt"

func main() {

	//Criou o arquivo
	file, err := os.Create(FILE)

	//Caso não seja possivel escrevere no arquivo dará um erro
	if err != nil {
		fmt.Println("Não foi possivel criar o arquivo...")
		return
	}

	defer file.Close()
	defer fmt.Println("Arquivo criado com sucesso.")

	languages := []string{"Java", "Kotlin", "Go", "Elixir"}

	for index, content := range languages {
		file.Write([]byte(content))

		if index-1 < len(languages) {
			file.Write([]byte("\n"))
		}
	}
}

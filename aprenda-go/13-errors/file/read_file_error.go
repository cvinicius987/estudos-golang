package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FILE = "aprenda-go/13-errors/file/languages.txt"

func main() {

	file, err := os.Open(FILE)

	if err != nil {
		fmt.Println("Não foi possivel abrir o arquivo...")
		return
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Não foi possivel ler os dados do arquivo")
		return
	}

	fmt.Println(string(bytes))
}

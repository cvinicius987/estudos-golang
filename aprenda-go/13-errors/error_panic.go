package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("names.txt")

	defer fmt.Println("Aqui defer...")

	//Executa os defers, mas não tem uma saida como o Panicln
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

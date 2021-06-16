package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer fmt.Println("Fim da execucao....")

	file, err := os.Open("names.txt")

	if err != nil {

		//Com esse as funções não serão executadas
		//O codigo de saida será exit 1
		//não executa os defers
		log.Fatalln(err)
	}

	defer file.Close()
}

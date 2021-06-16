package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("names.txt")

	defer fmt.Println("Aqui defer...")

	//Os defers serão executados, e será invocado o panic()
	//Encerra as funções na ordem
	if err != nil {
		log.Panicln("Aqui erro com Panicln...")
	}

	defer file.Close()
}

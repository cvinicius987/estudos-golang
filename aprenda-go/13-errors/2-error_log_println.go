package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Open("names.txt")

	if err != nil {
		log.Println("Aqui erro com log", err)
		return
	}

	defer file.Close()

}

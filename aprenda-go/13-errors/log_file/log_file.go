package main

import (
	"fmt"
	"log"
	"os"
)

const FILE = "aprenda-go/13-errors/log_file/log_file.txt"

func main() {

	//Arquivo de Log
	fileLog, _ := os.Create(FILE)

	defer fileLog.Close()

	log.SetOutput(fileLog)

	//Leitura do Processo
	file, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("Aqui erro com log", err)
	}

	defer file.Close()

	fmt.Println("Conclusao do arquivo de log")
}

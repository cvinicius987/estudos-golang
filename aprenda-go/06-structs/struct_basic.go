package main

import "fmt"

//Declaração de um struct
type language struct {
	name     string
	compiled bool
}

func main() {

	/*
	 * Uso completo de um struct
	 */
	language1 := language{
		name:     "Java",
		compiled: true,
	}

	language2 := language{
		name:     "Golang",
		compiled: true,
	}

	language3 := language{
		name:     "PHP",
		compiled: false,
	}

	//Uso de forma simplificado
	language4 := language{"Kotlin", true}

	fmt.Println(language1, language2, language3, language4)
}

package main

import "fmt"

func main() {

	//Ao criar um map sem inicializar devemos cria-lo com make()
	languages := make(map[string]int)

	languages["java"] = 1
	languages["kotlin"] = 2
	languages["golang"] = 3

	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}

}

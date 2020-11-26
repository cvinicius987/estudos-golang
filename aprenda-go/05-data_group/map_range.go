package main

import "fmt"

func main() {

	languages := map[string]int{
		"java":   1,
		"kotlin": 2,
		"golang": 3,
	}

	//Utilizando range
	//como não temos index a expressão retorna a key/value
	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}

}

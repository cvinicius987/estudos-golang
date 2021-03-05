package main

import "fmt"

func main() {

	languages := map[string]int{
		"java":   1,
		"kotlin": 2,
		"golang": 3,
	}

	fmt.Println(languages)

	//Realiza a exclusão de uma key do map
	delete(languages, "java")

	fmt.Println(languages)
}

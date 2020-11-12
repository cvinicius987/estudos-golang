package main

import "fmt"

func main() {

	name := "Google"

	lastname := "Golang"

	//É retornado uma nova String concatenando as encaminhadas como parametro
	result := fmt.Sprint(name, " ", lastname)

	fmt.Println(result)
}

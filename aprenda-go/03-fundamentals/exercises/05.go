package main

import "fmt"

/*
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
*/
func main() {

	name := `
		Golang
			é uma lingguagem 
				tipada e
					compilada
	`

	fmt.Println(name)
}

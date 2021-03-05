package main

import "fmt"

func main() {

	languages := []string{"Java", "Kotlin", "JavaScript", "Golang", "Elixir"}

	languages = append(languages[:2], languages[3:4]...)

	fmt.Println(languages)

}

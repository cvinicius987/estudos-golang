package main

import (
	"fmt"
	"sort"
)

func main() {

	list := []string{"Java", "Golang", "Kotlin", "Elixir", "C", "ABAP"}

	fmt.Println(list)

	sort.Strings(list)

	fmt.Println(list)
}

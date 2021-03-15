package main

import "fmt"

func main() {

	result := calculateVar(1, 2, 3, 4, 5)

	fmt.Println(result)
}

/**
  O parametro variadico deve ser sempre o ultimo parametro da função
*/
func calculateVar(x ...int) int {

	var result int

	for _, value := range x {
		result += value
	}

	return result
}

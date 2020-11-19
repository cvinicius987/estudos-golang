package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	//A função append concatena um slice com dados do mesmo tipo
	numbers_final := append(numbers, 6, 7, 8, 9, 10)

	fmt.Println(numbers_final)

	numbers2 := []int{11, 12, 13, 14, 15}

	//Para concatenar slices,
	//precisamos usar o operator ... que desdobra os valores de uma slice
	numbers_final = append(numbers_final, numbers2...)

	fmt.Println(numbers_final)
}

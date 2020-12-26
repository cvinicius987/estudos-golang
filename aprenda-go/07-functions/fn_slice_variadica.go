package main

import "fmt"

func main() {

	numbers := []int{1, 25, 36, 85, 69}

	result := calculateSlice(numbers...)

	//Função com parametros variadicos podem conter n elementos ou nenhum
	result2 := calculateSlice()

	fmt.Println(result, result2)
}

func calculateSlice(x ...int) int {

	var result int

	for _, value := range x {
		result += value
	}

	return result
}

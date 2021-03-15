package main

import "fmt"

//funções callback são funções que são passadas como argumentos
func main() {

	var numbers = []int{1, 2, 3, 4, 5}

	var impares = func(number int) bool {

		return number%2 != 0
	}

	var pares = func(number int) bool {

		return number%2 == 0
	}

	fmt.Println("Impares:", processCallback(somaCallback, impares, numbers))
	fmt.Println("Pares", processCallback(somaCallback, pares, numbers))
}

func somaCallback(numbers []int) int {

	var result = 0

	for _, value := range numbers {
		result += value
	}

	return result
}

func processCallback(calcule func(x []int) int, expression func(x int) bool, numbers []int) int {

	var temp []int

	for _, value := range numbers {

		if expression(value) {
			temp = append(temp, value)
		}
	}

	return calcule(temp)
}

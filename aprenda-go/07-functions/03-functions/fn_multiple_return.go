package main

import "fmt"

func main() {

	result, message := calculateMultiple("soma", 1, 2, 3, 4, 5)

	fmt.Println(result, message)
}

func calculateMultiple(operation string, x ...int) (int, string) {

	var message string
	var result int

	if operation == "soma" {

		for _, value := range x {
			result += value
		}

		message = "Soma"

	} else {

		for _, value := range x {
			result -= value
		}

		message = "Subtração"
	}

	return result, message
}

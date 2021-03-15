package main

import (
	"fmt"
)

/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/
func main(){

	fmt.Println(sumVar(1,2,3,4))

	fmt.Println(sumSlice([]int{1,2,3,4}))
}

func sumVar(numbers ...int) int{

	total := 0

	for _, value := range numbers{

		total += value
	}

	return total
}

func sumSlice(numbers []int) int{

	total := 0

	for _, value := range numbers{
		total += value
	}

	return total
}


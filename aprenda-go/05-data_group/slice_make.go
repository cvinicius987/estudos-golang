package main

import "fmt"

func main() {

	/**
	 * A função make serve para criar um slice com um tamanho inicial e uma capacidade
	 * extra, dessa forma quando adicionamos itens no slice o array interno não será refeito
	 */
	numbers := make([]int, 5, 10)

	numbers[0], numbers[1] = 1, 2
	numbers[2], numbers[3], numbers[4] = 3, 4, 5

	fmt.Println(numbers, "Tamanho:", len(numbers), "Capacidade:", cap(numbers))

	//Adiciona elementos no slice sem recriar o array, devido a sua capacidade
	//definido no make ser suficiente para atender a demanda inicial
	numbers = append(numbers, 6, 7, 8, 9, 10)

	fmt.Println(numbers, "Tamanho:", len(numbers), "Capacidade:", cap(numbers))

	//Cria um novo slice, e como foi feito através do make
	//define uma capacidade de armazenamento
	numbers = append(numbers, 11, 12, 13, 14, 15)

	fmt.Println(numbers, "Tamanho:", len(numbers), "Capacidade:", cap(numbers))
}

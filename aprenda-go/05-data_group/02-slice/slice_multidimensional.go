package main

import "fmt"

func main() {

	//Similar ao array multidimensionais de outras linguagens
	//Basicamente é um slice dentro de outro slice
	numbers := [][]int{
		[]int{1, 2, 3, 4, 5},

		[]int{6, 7, 8, 9, 10},

		[]int{11, 12, 13, 14, 15},

		[]int{16, 17, 18, 19, 20},
	}

	for index, obj := range numbers {

		for index2, obj2 := range obj {

			fmt.Println("Linha:", index, "Coluna:", index2, "Valor:", obj2)
		}
	}
}

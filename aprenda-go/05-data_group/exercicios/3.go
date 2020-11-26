package main

import "fmt"

/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	- Desafio: obtenha o mesmo resultado acima utilizando a função len()
	        para determinar o penúltimo item
*/

func main() {

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	fmt.Println(numbers[:3])

	//Do quinto ao último item do slice (incluindo o último item!)
	fmt.Println(numbers[4:])

	//Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println(numbers[2:9])

	//Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println(numbers[2:(len(numbers) - 1)])
}

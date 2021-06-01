package main

import "fmt"

/*
- Escreva um programa que coloque 100 números em um canal,
retire os números do canal, e demonstre-os.
*/
func main() {

	total := 100
	canal := make(chan int)

	go func(x int) {

		for i := 1; i <= x; i++ {
			canal <- i
		}

		close(canal)

	}(total)

	for v := range canal {
		fmt.Println("Valor Canal: ", v)
	}
}

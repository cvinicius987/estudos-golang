package main

import (
	"fmt"
)

/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/
func main() {

	channel := make(chan int)
	goNumbers := 10

	for i := 1; i <= goNumbers; i++ {

		go func() {

			for j := 1; j <= 10; j++ {
				channel <- j
			}
		}()
	}

	for i := 0; i < 100; i++ {

		fmt.Println(i, "\t", <-channel)
	}
}

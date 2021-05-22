package main

import "fmt"

func main() {

	channel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- i
		}
	}()

	for i := 1; i <= 10; i++ {

		v, ok := <-channel

		fmt.Println("Valor:", v, "Correto:", ok)
	}
}

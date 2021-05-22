package main

import "fmt"

func main() {

	//Declaração de um canal
	channel := make(chan int)

	//Envia valores para um canal
	//Aqui teremos erro pois uma mensagem deve ser encaminhada entre goroutines
	channel <- 40

	fmt.Println(<-channel)

}

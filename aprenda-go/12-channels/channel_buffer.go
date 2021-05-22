package main

import "fmt"

func main() {

	//Declaração de um canal com buffer valendo 2
	channel := make(chan int, 2)

	//Aqui devido ao buffer não teremos erro de block
	channel <- 51
	channel <- 52

	//Efetua a leitura dos dados que esta no canal
	fmt.Println(<-channel)
}

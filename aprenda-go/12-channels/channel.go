package main

import "fmt"

func main() {

	//Declaração de um canal
	//Canais podem ser bidirecionais: make(chan int)
	//				   send: make(chan<- int) (Envia dados ao canal)
	//				   received make(<-chan int) (Leitura de dados do canal)
	channel := make(chan int)

	//Envia valores para um canal
	//Esta goroutine envia dados para a goroutine Main
	go func() {
		channel <- 40
	}()

	//Efetua a leitura dos dados que esta no canal
	fmt.Println(<-channel)

}

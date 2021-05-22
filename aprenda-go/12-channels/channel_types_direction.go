package main

import "fmt"

/*
	Declaração de um canal
	Canais podem ser bidirecionais: make(chan int)
					   send: make(chan<- int) (Envia dados ao canal)
					   received make(<-chan int) (Leitura de dados do canal)

	Um canal pode ser declarado como bidirecional, send ou received;

	Podemos fazer conversões de canais, sempre em tipos compativeis:
	bidirecional pode ser um send ou received;

	Um send não pode ser um received;
*/
func main() {

	//Canal bidirecional
	channel := make(chan string)

	//Encaminha um canal bidirecional para uma função de somente leitura
	go write(channel)

	//Encaminha um canal bidirecional para uma função de somente escrita
	read(channel)

}

func read(channel <-chan string) {

	fmt.Println("A mensagem é: ", <-channel)
}

func write(channel chan<- string) {

	channel <- "Golang é uma linguagem concorrente..."
}

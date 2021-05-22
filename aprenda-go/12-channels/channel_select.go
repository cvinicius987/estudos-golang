package main

import "fmt"

/*
Select serve para ler dados de canais,
podemos ler dados de varios canais ao menos tempo usando o select

A Ordem de leitura acontece em qual canal que ter dados a serem lidos,
caso tenha dados em mais de um canal a leitura acontece de forma aleatoria
*/
func main() {

	number := 10

	channel1 := make(chan int)
	channel2 := make(chan int)

	go calcule(number, channel1)
	go calcule(number, channel2)

	for i := 0; i < (number * 2); i++ {

		select {
		case v := <-channel1:
			fmt.Println("Canal 1:", v)
		case v := <-channel2:
			fmt.Println("Canal 2:", v)
		}
	}
}

func calcule(num int, channel chan<- int) {

	for i := 0; i < num; i++ {

		channel <- i
	}
}

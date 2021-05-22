package main

import (
	"fmt"
	"sync"
)

var a sync.WaitGroup

func main() {

	channel := make(chan int)

	a.Add(2)

	go write(10, channel)
	go read(channel)

	a.Wait()
}

//Recebe um canal de escrita
func write(total int, c chan<- int) {

	for i := 0; i < total; i++ {
		c <- i
	}

	close(c)

	a.Done()
}

//Recebe um canal de leitura
func read(c <-chan int) {

	for content := range c {
		fmt.Println("Valor no canal:", content)
	}

	a.Done()
}

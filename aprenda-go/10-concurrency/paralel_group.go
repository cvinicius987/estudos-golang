package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
 Dessa forma declaração um sinalizador dizendo que vamos controlar a quantidade de Goroutines
 que será executada

 Um codigo só sera executado de forma paralela caso o ambientes tenha mais de um core,
 caso contrario mesmo que tenhamos um codigo que é concorrente, ele será executado de forma aleatoria
 e não em paralelo.
*/
var wg sync.WaitGroup

func main() {

	fmt.Println(" =========== Cores:", runtime.NumCPU())
	fmt.Println(" =========== Goroutines: ", runtime.NumGoroutine())

	//Avisa que o codigo terá mais 2 Goroutines
	wg.Add(2)

	go first()
	go second()

	fmt.Println(" =========== Goroutines: ", runtime.NumGoroutine())

	//Quando a goroutine Main chegar ao fim, manda esperar
	wg.Wait()
}

func first() {

	for i := 0; i < 100; i++ {
		fmt.Println("first: ", i)
		time.Sleep(10)
	}

	wg.Done()
}

func second() {

	for i := 0; i < 100; i++ {
		fmt.Println("second: ", i)
		time.Sleep(10)
	}

	wg.Done()
}

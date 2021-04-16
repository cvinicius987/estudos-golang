package main

import (
	"fmt"
	"runtime"
)

/*
	Neste exemplo a função first() não será executada, pois a goroutine main
	termina antes da execução da goroutine first()
*/
func main() {

	fmt.Println(" =========== Cores:", runtime.NumCPU())
	fmt.Println(" =========== Goroutines: ", runtime.NumGoroutine())

	go first()
	second()

	fmt.Println(" =========== Goroutines: ", runtime.NumGoroutine())
}

func first() {

	for i := 0; i < 10; i++ {
		fmt.Println("first: ", i)
	}
}

func second() {

	for i := 0; i < 10; i++ {
		fmt.Println("second: ", i)
	}
}

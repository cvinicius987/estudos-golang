package main

import (
	"fmt"
	"runtime"
)

/*
	Codigo executado em uma unica goroutine, sendo essa a main,
	codido será executado de forma linear.
*/
func main() {

	fmt.Println(" =========== Cores:", runtime.NumCPU())
	fmt.Println(" =========== Goroutines: ", runtime.NumGoroutine())

	first()
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

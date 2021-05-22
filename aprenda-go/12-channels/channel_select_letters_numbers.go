package main

import (
	"fmt"
	"unicode"
)

func main() {

	content := []int32{'a', '3', 'c', '1', '2', 'e', '7', 'f'}

	letters := make(chan int32)
	numbers := make(chan int32)

	quit := make(chan bool)

	go add(content, letters, numbers, quit)

	for {

		select {
		case v := <-letters:
			fmt.Printf("Temos a letra %s, binario: %b , Unicode: %#U\n", string(v), v, v)
		case v := <-numbers:
			fmt.Printf("Temos o numero %d, binario: %b , Unicode: %#U\n", v, v, v)
		case <-quit:
			return
		}
	}

}

func add(content []int32, letters chan<- int32, numbers chan<- int32, quit chan<- bool) {

	for _, c := range content {

		switch {
		case unicode.IsLetter(c):
			letters <- c
		case unicode.IsDigit(c):
			numbers <- c
		}
	}

	close(letters)
	close(numbers)

	quit <- true
}

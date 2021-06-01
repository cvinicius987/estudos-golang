package main

import (
	"fmt"
)

func main() {
	q := make(chan bool)
	c := gen(q)

	for {

		select {
		case v := <-c:
			fmt.Println("Canal C:", v)
		case <-q:
			fmt.Println("Recebeu Quit..........")
			return
		}
	}

}

func gen(q chan bool) <-chan int {
	c := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)

		q <- true
	}()

	close(q)

	return c
}

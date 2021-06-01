package main

import "fmt"

func main() {

	channel := make(chan int, 1)

	channel <- 10

	fmt.Println(<-channel)
}

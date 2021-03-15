package main

import "fmt"

func main() {

	x := 10

	//y tem como valor o ponteiro de x
	y := &x

	fmt.Println("x:", x, "y:", y)
	fmt.Printf("x type: %T, y type: %T\n", x, y)

	//incrementando o valor que esta localização no ponteiro
	*y++

	fmt.Println("x:", x, "y:", *y)
}

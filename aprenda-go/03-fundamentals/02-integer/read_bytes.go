package main

import "fmt"

func main() {

	x := "e"
	y := "ç"

	xB := []byte(x)
	yB := []byte(y)

	fmt.Printf("%v, %T, Bytes: %v\n", x, x, xB)
	fmt.Printf("%v, %T, Bytes: %v\n", y, y, yB)
}

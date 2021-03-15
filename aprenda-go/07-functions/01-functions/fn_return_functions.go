package main

import "fmt"

func main() {

	x := 10
	y := 10

	resultFn := calcule(x, y)

	fmt.Println(resultFn())
}

func calcule(x int, y int) func() int {

	result := x + y

	return func() int {
		return (result + 100)
	}
}

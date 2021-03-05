package main

import (
	"fmt"
)

/*
- Crie e utilize uma função anônima.
*/

func main(){

	func(x int, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
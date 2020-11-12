package main

import "fmt"

func main() {

	y := 35

	for x := 0; x <= 100; x++ {

		fmt.Println(x)

		if x == y {
			fmt.Println("Achou.....")
			break
		}
	}
}

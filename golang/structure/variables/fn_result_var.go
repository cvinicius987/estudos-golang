package main

import "fmt"

func main() {

	x, j := exec()

	fmt.Println(x, j)
}

func exec() (int, string) {

	return 1, "Name"
}

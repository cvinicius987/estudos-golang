package main

import "fmt"

var x interface{}

func main() {

	x = 10

	switch x.(type) {
	case bool:
		fmt.Println(x, " é bool")
	case string:
		fmt.Println(x, " é string")
	case int:
		fmt.Println(x, " é int")
	case float64:
		fmt.Println(x, " é float")
	}
}

package main

import (
	"fmt"
)

//Default: 0
var number int

//Default: 0
var money float64

//Defaut: ""
var name string

//Default: false
var isOk bool

func main() {

	fmt.Printf("%v, %T\n", number, number)
	fmt.Printf("%v, %T\n", money, money)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", isOk, isOk)
}

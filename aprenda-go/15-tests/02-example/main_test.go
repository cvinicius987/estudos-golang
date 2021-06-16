package main

import "fmt"

func ExampleSoma() {

	fmt.Println(Soma(10, 10))
	//Output: 20
}

func ExampleSoma_Many_Option() {

	fmt.Println(Soma(10, 10))
	fmt.Println(Soma(15, 20))
	fmt.Println(Soma(82, 100))
	//Output:
	//20
	//35
	//182
}

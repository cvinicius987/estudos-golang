package main

import "fmt"

/*
O iota permite declarar valores sequencias em um conjunto de constantes
*/
const (
	a = iota
	b = iota
	c = iota
)

/*
Caso desejamos pular um valor, basta usar o _ para desprezar uma constante
*/
const (
	d = iota + 10
	e
	_
	g
)

func main() {

	fmt.Println(a, b, c)
	fmt.Println(d, e, g)
}

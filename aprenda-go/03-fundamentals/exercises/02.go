package main

import "fmt"

/*
Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
==
!=
<=
<
=
*/

func main() {

	var a = 10

	b := a == 10

	c := a != 1000

	d := 1000 <= a

	e := 589 < a

	fmt.Println(a, b, c, d, e)

}

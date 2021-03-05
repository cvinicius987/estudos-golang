package main

import "fmt"

func main() {

	var i uint16 = 65535

	//Aqui temos erros de compilação devido a add um valor maior
	//que o limite de uint16
	//i = 65536

	fmt.Println(i)

	//Ao add na variavel que já armazena seu limite
	//ela é zerada e inicia do zero
	i++

	fmt.Println(i)

	i++

	fmt.Println(i)

}

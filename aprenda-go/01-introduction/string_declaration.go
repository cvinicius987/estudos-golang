package main

import "fmt"

func main() {

	//So pode ser usada em variaveis nivel de função
	s := "teste2"

	//Mais utilizada
	//Inicialização com valor default
	var s1 string

	//Usada quando declaramos multiplas variaveis
	var s2 = ""

	//Declaração valida quando atribuimos mais de uma variavei
	var s3 = "teste"

	//Declaração redundante
	var s4 string = ""

	//Declaração multipla
	s5, s6 := "teste_s5", "teste_s6"

	fmt.Println(s, s1, s2, s3, s4, s5, s6)
}

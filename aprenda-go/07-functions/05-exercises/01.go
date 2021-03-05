package main

/*
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
	- Demonstre seus resultados
*/
func main(){

	println(number())
	println(numberStr())
}

func number() int {

	return 10
}


func numberStr() (int, string) {

	return 10, "Golang"
}
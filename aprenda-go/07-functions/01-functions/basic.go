package main

import "fmt"

/*
Funções

- Funções são declaradas com parametros
- Funções são chamadas com argumentos
- Em Go os valores são passado por: pass by value

“Passing by value” refers to passing a copy of the value.
“Passing by reference” refers to passing the real reference of the variable in memory.
*/
func main() {
	value := soma(10, 20)

	fmt.Println(value)

	basic()
}

/*
Declaração curta em caso de serem do mesmo tipo

soma(x int, y int) ou soma(x int, y)
*/
func soma(x int, y int) int {
	return x + y
}

func basic() {
	fmt.Println("Aqui função basica.....")
}

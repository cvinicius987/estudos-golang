package main

import (
	"fmt"
)

/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/
func main(){

	result := operation("sum")

	fmt.Println(result(100, 20))
}

func operation(option string) func (x int, y int) int {

	switch option {
		case "minus":
			return func (x int, y int) int{
				return x - y
			}
		case "divide":
			return func (x int, y int) int{
				return x / y
			}
		default:
			return func (x int, y int) int{
				return x + y
			}
	}
}
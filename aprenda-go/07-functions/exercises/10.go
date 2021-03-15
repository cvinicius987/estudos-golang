package main

import "fmt"
/*
 Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, 
	onde esta outra função faz uso de uma variável alem de seu scope.
*/

func main(){
	
	discount := 10.0

	fn := sell(700.00, discount)

	fmt.Println(fn())
}

func sell(value float64, discount float64) func() float64{

	return func() float64{

		result :=  value - (value / 100 * discount)

		return result
	}
}
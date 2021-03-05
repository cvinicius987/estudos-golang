package main

import "fmt"

/*
 Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
    - Package-level scope
    - Function-level scope
    - Code-block-in-code-block scope
- Closures nos permitem salvar dados entre function calls e 
  ao mesmo tempo isolar estes dados do resto do código.
*/

func main(){

	//Cada chamada de increment retorna uma escopo de função
	callOne := increment()

	fmt.Println(callOne())
	fmt.Println(callOne())
	fmt.Println(callOne())

	fmt.Println("\n")

	callTwo := increment()

	fmt.Println(callTwo())
	fmt.Println(callTwo())
	fmt.Println(callTwo())
}

func increment() func() int{

	temp := 0

	return func() int{
		temp++
		return temp
	}
}


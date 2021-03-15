package main

import(
	"math/rand"
	"fmt"
)
/*
- Callback: passe uma função como argumento a outra função.
*/

func main(){
	
	fn := func(x int, y int) int{
		return x + y
	}

	process(fn)
}

func process(paramFn func (x int, y int) int){

	fmt.Println(paramFn(rand.Intn(100), rand.Intn(100)))
}
package main

import (
	"fmt"
)

/*
- Atribua uma função a uma variável.
- Chame essa função.
*/
func main(){

	showName := func(name string, middleName string){
		
		fmt.Println(name, middleName)
	}

	showName("Go", "Lang")

}
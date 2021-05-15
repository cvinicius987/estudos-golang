package main

import (
	"fmt"
	"sync"
)

/*
Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines
finalizem antes de o programa terminar.
*/
var w sync.WaitGroup

func main() {

	w.Add(2)

	go print("Aqui 1....")
	go print("Aqui 2....")

	w.Wait()

	fmt.Println("Main...")
}

func print(msg string) {
	fmt.Println("A mensagem é: ", msg)

	w.Done()
}

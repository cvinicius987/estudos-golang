package main

import "fmt"

/**
- O defer adia a execução de uma determinada função
- Em caso de multiplos defer o que entrou primeiro será executado por ultimo
- o Defer deve ser executado entes do encerramento que originou o defer
*/
func main() {

	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}

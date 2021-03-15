package main

//Presente dentro do $GOPATH, podemos usar alias para o nome do package
import (
	"fmt"
	calc "go-in-action-examples/3-packages-tools/calculator"
)

func main() {

	fmt.Println(calc.Sum(10, 20))
}

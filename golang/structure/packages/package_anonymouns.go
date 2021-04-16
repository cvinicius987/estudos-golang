package main

import (
	"fmt"
	_ "go-in-action-examples/3-packages-tools/driver" //_ significa que não usara nada do package
)

func main() {
	fmt.Println("Aqui sem usar o conteudo do driver")
}

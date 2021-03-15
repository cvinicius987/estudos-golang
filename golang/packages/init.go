package main

/**
* A função chamada init dentro de um package é chamada
  ao carregar o package
*/
import (
	dri "go-in-action-examples/3-packages-tools/driver"
)

func main() {
	dri.Open()
}

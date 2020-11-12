package main

import "fmt"

func main() {

	for horas := 0; horas <= 12; horas++ {

		for minutos := 0; minutos <= 59; minutos++ {

			for segundos := 0; segundos <= 59; segundos++ {

				fmt.Println(horas, ":", minutos, ":", segundos)
			}
		}
	}

}

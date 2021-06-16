package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("Aqui erro com println...")
		return
	}

	defer file.Close()

}

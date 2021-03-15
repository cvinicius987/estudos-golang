package main

import "fmt"

func main() {

	list := []string{"Java", "Kotlin", "Go"}

	toMap := func(items []string) map[int]string {

		var mapT = map[int]string{}

		for index, value := range items {
			mapT[index] = value
		}

		return mapT
	}

	fmt.Println(toMap(list))
}

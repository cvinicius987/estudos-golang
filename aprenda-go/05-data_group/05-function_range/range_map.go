package main

import "fmt"

func main() {

	languages := map[string]int{
		"Java":       1,
		"Kotlin":     2,
		"JavaScript": 3,
		"Golang":     4,
	}

	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %d, Value access map: %d \n", key, value, languages[key])
	}
}

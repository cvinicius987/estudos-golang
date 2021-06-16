package main

import "fmt"

func main() {

	valor := Average(8, 8, 8, 10)

	fmt.Println(valor)
}

func Average(notes ...float64) float64 {

	sum := 0.0

	for _, v := range notes {
		sum += v
	}

	return sum / float64(len(notes))
}

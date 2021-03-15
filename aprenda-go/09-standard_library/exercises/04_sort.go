package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	sort.Slice(xi, func(i, j int) bool {
		return xi[i] < xi[j]
	})

	fmt.Println("Xi Ascendente:", xi)

	sort.Slice(xi, func(i, j int) bool {
		return xi[i] > xi[j]
	})

	fmt.Println("Xi Decrescente:", xi)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})

	fmt.Println("xs Ascendente: ", xs)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] > xs[j]
	})

	fmt.Println("xs Decrescente", xs)
}

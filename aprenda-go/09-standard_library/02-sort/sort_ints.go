package main

import (
	"fmt"
	"sort"
)

func main() {

	list := []int{1, 5, 2, 3, 9, 78, 65, 545, 78,
		36, 45, 25, 20, 14, 144, 78520,
		2545582, 254520, 5}

	fmt.Println(list)

	sort.Ints(list)

	fmt.Println(list)
}

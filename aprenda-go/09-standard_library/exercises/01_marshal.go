package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{"Moneypenny", 27}

	result1, _ := json.Marshal(u1)

	result2, _ := json.Marshal(u2)

	fmt.Println(string(result1))

	fmt.Println(string(result2))
}

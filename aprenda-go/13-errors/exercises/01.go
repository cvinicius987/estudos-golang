package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

/*
- ...remova o underscore e verifique e lide com o erro de maneira apropriada.
*/
func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Println("Erro ao converter em JSON")
	}

	fmt.Println(string(bs))
}

package main

import (
	"encoding/json"
	"fmt"
)

type Cnpj string

type Address struct {
	Log    string
	Number int
}

type Client struct {
	Name    string
	Doc     Cnpj `json:"document"`
	Address Address
}

func main() {

	var client1 string = `{"Name":"Google","document":"123456789","Address":{"Log":"Avenue","Number":100}}`

	var result Client

	json.Unmarshal([]byte(client1), &result)

	fmt.Println(result)
}

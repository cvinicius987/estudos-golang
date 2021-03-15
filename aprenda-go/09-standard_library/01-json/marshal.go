package main

import (
	"encoding/json"
	"fmt"
)

/*
 * Os objetos para serem vistos fora do package devem ter a Inicial como UPPERCASE
 *
 * Para serializar um struct para JSON os atributos devem começar em UPPERCASE
 * https://cdn.rawgit.com/GoesToEleven/golang-web-dev/17e3852d/040_json/README.html
 */
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

	client1 := Client{
		Name:    "Google",
		Doc:     "123456789",
		Address: Address{"Avenue", 100},
	}

	result, err := json.Marshal(client1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Tamanho em Bytes: ", len(result))
	fmt.Println(result)
	fmt.Println(string(result))
}

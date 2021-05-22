package main

import "fmt"

type Client interface {
	GetDocument() string
}

type ClientPF struct {
	Id  int
	Cpf string
}

type ClientPJ struct {
	Id   int
	Cnpj string
}

func (clientPf ClientPF) GetDocument() string {
	return fmt.Sprintf("Documento é CPF: %s", clientPf.Cpf)
}

func (clientPj ClientPJ) GetDocument() string {
	return fmt.Sprintf("Documento é CNPJ: %s", clientPj.Cnpj)
}

func main() {

	clientPF := ClientPF{1, "123.456.789/20"}

	clientPJ := ClientPJ{2, "123.456.789.0001-20"}

	detailDocument(clientPF)
	detailDocument(clientPJ)
}

func detailDocument(client Client) {

	fmt.Println(client.GetDocument())
}

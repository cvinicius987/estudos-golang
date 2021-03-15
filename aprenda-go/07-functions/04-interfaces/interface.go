package main

import "fmt"

/**
Em Go não precisamos implementar uma interface diretamente
*/
type cnpj string
type cpf string

type Pessoa interface {
	mensagemPessoa()
}

type PessoaJuridica struct {
	name string
	cnpj
}

func (pj PessoaJuridica) mensagemPessoa() {
	fmt.Println("Ola", pj.name, "você é uma empresa com CNPJ", pj.cnpj)
}

type PessoaFisica struct {
	name string
	cpf
}

func (pf PessoaFisica) mensagemPessoa() {
	fmt.Println("Ola", pf.name, "você é uma pessoa com CPF", pf.cpf)
}

func main() {

	pj := PessoaJuridica{
		name: "Google Inc.",
		cnpj: "51.840.334/0001-04",
	}

	pf := PessoaFisica{
		name: "Caio",
		cpf:  "701.877.740-21",
	}

	emiteNotaFiscal(pj)
	fmt.Print("\n")
	emiteNotaFiscal(pf)
}

func emiteNotaFiscal(param Pessoa) {

	param.mensagemPessoa()

	switch param.(type) {
	case PessoaFisica:
		fmt.Println("Você é pessoa fisica não pode emitir nota fiscal.")
	case PessoaJuridica:
		fmt.Println("Nota Fiscal emitida com sucesso.")
	default:
		fmt.Println("Seu tipo não esta definido")
	}
}

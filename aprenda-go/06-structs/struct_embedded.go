package main

import "fmt"

type category struct {
	name string
}

type product struct {
	name  string
	value float64
	category
}

func main() {

	category1 := category{
		name: "Eletronicos",
	}

	category2 := category{
		name: "Eletrodomesticos",
	}

	//Declaracão normal
	product1 := product{
		name:     "TV",
		value:    2500.00,
		category: category1,
	}

	//Declaracão curta
	product2 := product{"Geladeira", 4752.15, category2}

	//Criando de forma embarcada
	product3 := product{
		name:  "Toalha",
		value: 10.55,
		category: category{
			name: "Cama, Mesa e Banho",
		},
	}

	fmt.Println(product1, product2, product3)
}

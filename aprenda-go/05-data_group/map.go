package main

import "fmt"

func main() {

	languages := map[string]int{
		"java":   1,
		"kotlin": 2,
		"golang": 3,
	}

	fmt.Println(languages, "\n")

	//Add Item
	languages["elixir"] = 4

	fmt.Println(languages, "\n")

	fmt.Println(languages["java"], "\n")

	//Ao chamar uma Key que não existe retorna 0
	fmt.Println(languages["php"], "\n")

	//O vAlor de retorno é zero mas podemos ter acesso a uma variavel que fala se o item existe
	resultPHP, existe := languages["php"]
	fmt.Println("Existe:", existe, "Valor", resultPHP)

	//utilizando if para testar o valor
	if conteudo, ok := languages["javascript"]; ok {
		fmt.Println("Existe o valor", conteudo)
	} else {
		fmt.Println("Não existe")
	}
}

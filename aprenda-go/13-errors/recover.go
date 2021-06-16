package main

import "fmt"

func main() {

	//Chamada da função f
	f()

	fmt.Println("Processo sendo retornado.....")
}

func f() {

	//Registra um defer antes da chamada a função g que irá lancar um panic()
	defer func() {

		//Ao executar esse defer, o recover avisa que deve se recuperar do erro
		//o o resto do programa continuara sendo executado
		if r := recover(); r != nil {
			fmt.Println("Se recuperando do erro...")
		}
	}()

	g()
}

func g() {

	for i := 0; i < 5; i++ {

		fmt.Printf("Item de g %d \n", i)

		//Quando chegar em 3 deve lancar um panic
		if i == 3 {
			panic("Aqui 3 paniccccccccc.")
		}
	}
}

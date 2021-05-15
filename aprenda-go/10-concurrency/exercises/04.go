package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
 Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/

var w sync.WaitGroup

func main() {

	var times = 1000
	var count = 0

	w.Add(times)

	var mu sync.Mutex

	for i := 1; i <= times; i++ {

		go func() {

			mu.Lock()

			temp := count

			runtime.Gosched()

			temp++

			count = temp

			w.Done()

			mu.Unlock()

		}()

	}

	w.Wait()

	fmt.Println("Esperado:", "1000", "Total:", count)
}

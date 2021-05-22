package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
  Utilize atomic para consertar a condição de corrida do exercício #3.
*/

var w sync.WaitGroup

func main() {

	var times = 1000
	var count int32 = 0

	w.Add(times)

	for i := 1; i <= times; i++ {

		go func() {

			atomic.AddInt32(&count, 1)

			runtime.Gosched()

			w.Done()

		}()

	}

	w.Wait()

	fmt.Println("Esperado:", "1000", "Total:", count)
}

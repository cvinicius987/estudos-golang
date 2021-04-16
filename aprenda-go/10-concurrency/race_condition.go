package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	count := 0
	processes := 1000

	var w sync.WaitGroup

	w.Add(processes)

	for i := 0; i < processes; i++ {

		go func() {

			v := count

			runtime.Gosched()

			v++

			count = v

			w.Done()
		}()
	}

	w.Wait()
	fmt.Println("Resultado:", count)
}

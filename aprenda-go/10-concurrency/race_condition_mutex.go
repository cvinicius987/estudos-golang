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

	var mu sync.Mutex

	for i := 0; i < processes; i++ {

		go func() {

			mu.Lock()

			v := count

			runtime.Gosched()

			v++

			count = v

			w.Done()

			mu.Unlock()
		}()
	}

	w.Wait()
	fmt.Println("Resultado:", count)
}

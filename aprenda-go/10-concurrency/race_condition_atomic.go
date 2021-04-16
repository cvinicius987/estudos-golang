package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var count int64 = 0
	processes := 1000

	var w sync.WaitGroup

	w.Add(processes)

	for i := 0; i < processes; i++ {

		go func() {

			atomic.AddInt64(&count, 1)

			runtime.Gosched()

			w.Done()
		}()
	}

	w.Wait()
	fmt.Println("Resultado:", count)
}

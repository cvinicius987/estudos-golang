package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("Arch: %s \nOs: %s \n ", runtime.GOARCH, runtime.GOOS)
}

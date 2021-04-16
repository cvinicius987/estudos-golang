package main

import "fmt"

func main() {

	count := 0

	v := count

	v++

	fmt.Println("count: ", count, &count)
	fmt.Println("v: ", v, &v)

}

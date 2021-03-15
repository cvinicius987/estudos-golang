package main

import "fmt"

/*
Os tipos das constantes são definidos no momento de seu uso,
diferente do tipo var
*/
const address = "127.0.0.1"

const (
	ip   = "172.0.0.1"
	port = 8080
)

func main() {

	fmt.Printf("%v, %T\n", address, address)
	fmt.Printf("%v, %T\n", ip, ip)
	fmt.Printf("%v, %T\n", port, port)
}

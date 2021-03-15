package main

import "fmt"

type client struct {
	name  string
	value int
}

func main() {

	client1 := client{"Google", 1}

	incrementCopyFromClient(client1)
	incrementReferenceFromClient(&client1)

	fmt.Println("Main:", client1)
}

func incrementCopyFromClient(c client) {

	c.value++

	fmt.Println("Increment Copy:", c)
}

func incrementReferenceFromClient(c *client) {

	c.value++
	(*c).value++

	fmt.Println("Increment Reference:", *c)
}

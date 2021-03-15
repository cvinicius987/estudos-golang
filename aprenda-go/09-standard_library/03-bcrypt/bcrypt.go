package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pass := "20julho1980"
	invalidPass := "20julho1980"

	sb, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bytes:", sb)
	fmt.Println("String:", string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(pass)))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(invalidPass)))
}

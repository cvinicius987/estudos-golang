package main

import "fmt"

func main() {

	//1 Byte
	var x8 int8 = 127

	//2 Bytes
	var x16 int16 = 10

	//4 Bytes
	var x32 int32 = 10

	//8 Bytes
	var x64 int64 = 10

	fmt.Println(" ============= Int ===============")
	fmt.Printf("%v, %T\n", x8, x8)
	fmt.Printf("%v, %T\n", x16, x16)
	fmt.Printf("%v, %T\n", x32, x32)
	fmt.Printf("%v, %T\n", x64, x64)

	fmt.Println(" ============= UInt ===============")

	/*
		O tipos Unsigned só possuem numeros positivos e não podem
		conter numeros negativos
	*/

	//1 Byte
	var y8 uint8 = 127

	//2 Bytes
	var y16 uint16 = 11452

	//4 Bytes
	var y32 uint32 = 14525485

	//8 Bytes
	var y64 uint64 = 1145254525

	fmt.Printf("%v, %T\n", y8, y8)
	fmt.Printf("%v, %T\n", y16, y16)
	fmt.Printf("%v, %T\n", y32, y32)
	fmt.Printf("%v, %T\n", y64, y64)

}

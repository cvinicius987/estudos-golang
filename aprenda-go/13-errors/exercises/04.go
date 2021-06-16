package main

import (
	"errors"
	"fmt"
	"log"
)

/*
...use o struct sqrt.Error como valor do tipo erro.
*/

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {

	var error = sqrtError{"Latitude erro", "Longitude erro", errors.New("norgate math: square root of negative number")}

	if f < 0 {
		return 0, error
	}

	return 42, nil
}

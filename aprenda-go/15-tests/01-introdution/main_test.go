package main

import (
	"log"
	"testing"
)

func TestAverage_CalculatedSuccess(t *testing.T) {

	valor := Average(8, 8.5, 9.5, 6)

	if valor != 8 {
		t.Error("Value:", valor, "Expected:", 8)
	}
}

func TestAverage_CalculatedFail(t *testing.T) {

	valor := Average(8, 8.5, 10, 6)

	if valor != 8 {
		t.Error("Value:", valor, "Expected:", 8)
	}
}

type testArr struct {
	data   []float64
	result float64
}

func TestAverage_CalculatedWithTable(t *testing.T) {

	table := []testArr{
		testArr{[]float64{4, 5, 6, 7, 8, 9, 10}, 7},
		testArr{[]float64{10, 12, 5, 12, 13, 11}, 10.5},
		testArr{[]float64{7, 8, 6, 7}, 7},
	}

	for _, v := range table {

		valor := Average(v.data...)

		log.Println("Value:", valor, "Expected:", v.result)

		if valor != v.result {
			t.Error("Value:", valor, "Expected:", v.result)
		}
	}
}

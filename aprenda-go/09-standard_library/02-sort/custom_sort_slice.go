package main

import (
	"fmt"
	"sort"
)

type employe struct {
	name   string
	age    int
	salary float64
}

func main() {

	list := []employe{
		{name: "Caio", age: 30, salary: 15000.00},
		{name: "Maria", age: 45, salary: 30000.48},
		{name: "Arnaldo", age: 35, salary: 12000.51},
		{name: "Rosana", age: 22, salary: 8000.00},
	}

	fmt.Println(list)

	sort.Slice(list, func(i, j int) bool {
		return list[i].name < list[j].name
	})

	fmt.Println("Ordenação por Name:", list)

	sort.Slice(list, func(i, j int) bool {
		return list[i].age < list[j].age
	})

	fmt.Println("Ordenação por Age:", list)

	sort.Slice(list, func(i, j int) bool {
		return list[i].salary < list[j].salary
	})

	fmt.Println("Ordenação por Salary:", list)
}

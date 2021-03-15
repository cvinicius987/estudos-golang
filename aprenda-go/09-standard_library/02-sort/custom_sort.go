package main

import (
	"fmt"
	"sort"
)

type car struct {
	name   string
	engine float64
}

//=============== Ordenação por Engine ascendente
type sortByEngine []car

func (c sortByEngine) Len() int {
	return len(c)
}

func (c sortByEngine) Less(i, j int) bool {
	return c[i].engine < c[j].engine
}

func (c sortByEngine) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//============= Ordenação por Engine decrescente
type sortByEngineDesc []car

func (c sortByEngineDesc) Len() int {
	return len(c)
}

func (c sortByEngineDesc) Less(i, j int) bool {
	return c[i].engine > c[j].engine
}

func (c sortByEngineDesc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//================ Ordenação por Nome
type sortByName []car

func (c sortByName) Len() int {
	return len(c)
}

func (c sortByName) Less(i, j int) bool {
	return c[i].name < c[j].name
}

func (c sortByName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {

	list := []car{
		{"Marea", 2.4},
		{"Monza", 2.0},
		{"HB20", 1.6},
		{"Palio", 1.0},
		{"Omega", 3.0},
	}

	fmt.Println("Sem ordenação:", list)

	sort.Sort(sortByEngine(list))

	fmt.Println("Com Ordenação Engine ASC:", list)

	sort.Sort(sortByEngineDesc(list))

	fmt.Println("Com Ordenação Engine Desc:", list)

	sort.Sort(sortByName(list))

	fmt.Println("Com Ordenação Name ASC:", list)
}

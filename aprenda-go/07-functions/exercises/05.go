package main

import (
	"fmt"
	"math"
)

/*
Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/

type figura interface{
	area() float64
}

type quadrado struct{
	lado float64 
}

type circulo struct{
	raio float64 
}

func (q quadrado) area() float64{
	return q.lado * q.lado
}

func (c circulo) area() float64{
	return math.Pi * 2 * c.raio
}

func main(){

	circulo1 := circulo{
		raio: 3.0,
	}

	quadrado1 := quadrado{
		lado: 2.0,
	}

	fmt.Println("Circulo: ", info(circulo1))
	fmt.Println("Quadrado: ", info(quadrado1))
}

func info(f figura) float64{

	return f.area()
}
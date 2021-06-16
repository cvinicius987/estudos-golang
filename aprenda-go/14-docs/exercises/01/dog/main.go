//dog - conteudo referente as funções de logica sobre regra caninas
package dog

import "log"

//fixAgeDog Padrão de anos dos cachorros
const fixAgeDog = 7

//Age com base no parametro age (idade humana) calcula a idade canina
func Age(age int) int {

	if age <= 0 {
		log.Fatalln("Não é possivel calcular a idade")
	}

	return age * fixAgeDog
}

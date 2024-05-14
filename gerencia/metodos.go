package gerencia

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PauloPHAL/Trabalho-AgeOfMitology-Go/domain"
)

func Jogar(lado1, lado2 *[]domain.Lutador) {
	somarPesos(lado1, lado2)
	maisVelho(lado1, lado2)

	rand.Seed(time.Now().UnixNano())
	if len(*lado1) > 0 && len(*lado2) > 0 {
		aleatorio := rand.Intn(20) + 1

		guerreiro1 := (*lado1)[0]
		guerreiro2 := (*lado2)[0]

		verQuemAtaca(aleatorio, guerreiro1, guerreiro2, lado1, lado2)
	}
}

func verQuemAtaca(aleatorio int, guerreiro1, guerreiro2 domain.Lutador, lado1, lado2 *[]domain.Lutador) {
	var permissao bool

	if aleatorio < 10 {
		permissao = atacarCadaUm(guerreiro1, guerreiro2, lado1, lado2)
		if permissao {
			atacarCadaUm(guerreiro2, guerreiro1, lado1, lado2)
		}
	} else {
		atacarCadaUm(guerreiro2, guerreiro1, lado1, lado2)
		atacarCadaUm(guerreiro1, guerreiro2, lado1, lado2)
	}
}

func atacarCadaUm(guerreiroAtacante, guerreiroSofredor domain.Lutador, lado1, lado2 *[]domain.Lutador) bool {
	codigo := bool(true)
	if guerreiroAtacante.GetEnergia() > 0 {
		codigo = guerreiroAtacante.Atacar(&guerreiroSofredor, lado1, lado2)
	}
	return codigo
}

func somarPesos(lado1, lado2 *[]domain.Lutador) {
	var pesoLado1, pesoLado2 int

	for _, lutador := range *lado1 {
		pesoLado1 += lutador.GetPeso()
	}

	for _, lutador := range *lado2 {
		pesoLado2 += lutador.GetPeso()
	}

	fmt.Println("Gregos e Nordicos pesam:", pesoLado1, "Kg")
	fmt.Println("Atlantes e Egipcios pesam:", pesoLado2, "Kg")
}

func maisVelho(lado1, lado2 *[]domain.Lutador) {
	var maiorIdade, aux = 0, 0

	for _, lutador := range *lado1 {
		if lutador.GetIdade() > maiorIdade {
			aux = lutador.GetIdade()
			if aux > maiorIdade {
				maiorIdade = aux
			}
		}
	}

	for _, lutador := range *lado2 {
		if lutador.GetIdade() > maiorIdade {
			aux = lutador.GetIdade()
			if aux > maiorIdade {
				maiorIdade = aux
			}
		}
	}

	fmt.Println("O mais velho tem", maiorIdade, "anos")
}

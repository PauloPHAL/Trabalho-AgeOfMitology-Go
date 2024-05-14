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

	for len(*lado1) > 0 && len(*lado2) > 0 {
		aleatorio := rand.Intn(20) + 1

		guerreiro1 := (*lado1)[0]
		guerreiro2 := (*lado2)[0]

		verQuemAtaca(aleatorio, guerreiro1, guerreiro2, lado1, lado2)

		moverGuerreiroDaFilaEVerificar(guerreiro1, lado1)
		moverGuerreiroDaFilaEVerificar(guerreiro2, lado2)
	}
	verificarVencedor(lado1, lado2)
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

func verificarVencedor(fila1, fila2 *[]domain.Lutador) {
	if len((*fila1)) > 0 {
		fmt.Println("Gregos e Nórdicos venceram.")
		fmt.Println()
		fmt.Println("GUERREIRO VITORIOSO: DADOS")
		guerreiroVitorioso := (*fila1)[len((*fila1))-1]
		fmt.Println("NOME:", guerreiroVitorioso.GetNome())
		fmt.Println("IDADE:", guerreiroVitorioso.GetIdade())
		fmt.Println("PESO:", guerreiroVitorioso.GetPeso())
	}
	if len((*fila2)) > 0 {
		fmt.Println("Atlantes e Egípcios venceram.")
		fmt.Println()
		fmt.Println("GUERREIRO VITORIOSO: DADOS")
		guerreiroVitorioso := (*fila2)[len((*fila2))-1]
		fmt.Println("NOME:", guerreiroVitorioso.GetNome())
		fmt.Println("IDADE:", guerreiroVitorioso.GetIdade())
		fmt.Println("PESO:", guerreiroVitorioso.GetPeso())
	}
}

func moverGuerreiroDaFilaEVerificar(guerreiro domain.Lutador, fila *[]domain.Lutador) {
	for i, g := range *fila {
		if g == guerreiro {
			*fila = append((*fila)[:i], (*fila)[i+1:]...)
			break
		}
	}

	if guerreiro.GetEnergia() > 0 {
		*fila = append(*fila, guerreiro)
	}

	if len(*fila) == 0 {
		fmt.Println("ULTIMO GUERREIO MORTO: DADOS")
		fmt.Println()
		fmt.Println("NOME: ", guerreiro.GetNome())
		fmt.Println("IDADE: ", guerreiro.GetIdade())
		fmt.Println("PESO: ", guerreiro.GetPeso())
		fmt.Println()
	}

	verificarLista(fila)
}

func verificarLista(fila *[]domain.Lutador) {
	for i := 0; i < len(*fila); i++ {
		if (*fila)[i].GetEnergia() <= 0 {
			*fila = append((*fila)[:i], (*fila)[i+1:]...)
			i = 0
		}
	}
}

package domain

import "fmt"

type Prometeano struct {
	Atlantes
}

func NewPrometeano(nome string, idade int, peso int, energia int) *Prometeano {
	return &Prometeano{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (p *Prometeano) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Prometeano atacou")
}

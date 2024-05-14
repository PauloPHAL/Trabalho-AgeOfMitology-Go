package domain

import "fmt"

type Prometeano struct {
	Atlantes
}

func NewPrometeano(nome string, idade int, peso int, energia int) *Prometeano {
	return &Prometeano{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (p *Prometeano) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Prometeano atacou")
	return true
}

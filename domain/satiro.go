package domain

import "fmt"

type Satiro struct {
	Atlantes
}

func NewSatiro(nome string, idade int, peso int, energia int) *Satiro {
	return &Satiro{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (s *Satiro) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Satiro atacou")
}

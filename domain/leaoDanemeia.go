package domain

import "fmt"

type LeaoDaNemeia struct {
	Gregos
}

func NewLeaoDaNemeia(nome string, idade int, peso int, energia int) *LeaoDaNemeia {
	return &LeaoDaNemeia{Gregos{Guerreiro{nome, idade, peso, energia}}}
}

func (l *LeaoDaNemeia) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("LeaoDaNemeia atacou")
	return true
}

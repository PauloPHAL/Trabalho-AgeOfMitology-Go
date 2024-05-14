package domain

import "fmt"

type Argus struct {
	Atlantes
}

func NewArgus(nome string, idade int, peso int, energia int) *Argus {
	return &Argus{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Argus) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Argus atacou")
	return true
}

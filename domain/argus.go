package domain

import "fmt"

type Argus struct {
	Atlantes
}

func NewArgus(nome string, idade int, peso int, energia int) *Argus {
	return &Argus{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Argus) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Argus atacou")
}

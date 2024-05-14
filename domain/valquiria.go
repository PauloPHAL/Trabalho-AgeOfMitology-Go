package domain

import "fmt"

type Valquiria struct {
	Nordicos
}

func NewValquiria(nome string, idade int, peso int, energia int) *Valquiria {
	return &Valquiria{Nordicos{Guerreiro{nome, idade, peso, energia}}}
}

func (v *Valquiria) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Valquiria atacou")
	return true
}

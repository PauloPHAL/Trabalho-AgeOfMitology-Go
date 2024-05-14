package domain

import "fmt"

type Anubita struct {
	Egipcios
}

func NewAnubita(nome string, idade int, peso int, energia int) *Anubita {
	return &Anubita{Egipcios{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Anubita) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Anubita atacou")
}

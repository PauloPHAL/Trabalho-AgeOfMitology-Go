package domain

import "fmt"

type Anubita struct {
	Egipcios
}

func NewAnubita(nome string, idade int, peso int, energia int) *Anubita {
	return &Anubita{Egipcios{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Anubita) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Anubita atacou")
	return true
}

package domain

import "fmt"

type Mumia struct {
	Egipcios
}

func NewMumia(nome string, idade int, peso int, energia int) *Mumia {
	return &Mumia{Egipcios{Guerreiro{nome, idade, peso, energia}}}
}

func (m *Mumia) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Mumia atacou")
	return true
}

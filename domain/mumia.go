package domain

import "fmt"

type Mumia struct {
	Egipcios
}

func NewMumia(nome string, idade int, peso int, energia int) *Mumia {
	return &Mumia{Egipcios{Guerreiro{nome, idade, peso, energia}}}
}

func (m *Mumia) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Mumia atacou")
}

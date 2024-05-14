package domain

import (
	"fmt"
)

type Ciclope struct {
	Gregos
}

func NewCiclope(nome string, idade, peso, energia int) *Ciclope {
	return &Ciclope{Gregos{Guerreiro{Nome: nome, Idade: idade, Peso: peso, Energia: energia}}}
}

func (c *Ciclope) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Ciclope atacou")
	return true
}

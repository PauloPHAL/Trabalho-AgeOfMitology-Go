package domain

import (
	"fmt"
)

type Ciclope struct {
	Gregos
}

func NewCiclope(nome string, idade, peso, energia int) *Ciclope {
	return &Ciclope{
		Gregos: Gregos{
			Guerreiro: Guerreiro{
				Nome:    nome,
				Idade:   idade,
				Peso:    peso,
				Energia: energia,
			},
		},
	}
}

func (c *Ciclope) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("Ciclope atacou")
}

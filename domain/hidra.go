package domain

import "fmt"

type Hidra struct {
	Gregos
}

func NewHidra(nome string, idade int, peso int, energia int) *Hidra {
	return &Hidra{Gregos{Guerreiro{nome, idade, peso, energia}}}
}

func (h *Hidra) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("Hidra atacou")
	return true
}

package domain

import "fmt"

type LoboDeFenris struct {
	Nordicos
}

func NewLoboDeFenris(nome string, idade int, peso int, energia int) *LoboDeFenris {
	return &LoboDeFenris{Nordicos{Guerreiro{nome, idade, peso, energia}}}
}

func (l *LoboDeFenris) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	fmt.Println("LoboDeFenris atacou")
	return true
}

package domain

import "fmt"

type LoboDeFenris struct {
	Nordicos
}

func NewLoboDeFenris(nome string, idade int, peso int, energia int) *LoboDeFenris {
	return &LoboDeFenris{Nordicos{Guerreiro{nome, idade, peso, energia}}}
}

func (l *LoboDeFenris) Atacar(lado1, lado2 *[]Lutador) {
	fmt.Println("LoboDeFenris atacou")
}

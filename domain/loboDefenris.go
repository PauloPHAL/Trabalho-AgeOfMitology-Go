package domain

type LoboDeFenris struct {
	Nordicos
}

func NewLoboDeFenris(nome string, idade int, peso int, energia int) *LoboDeFenris {
	return &LoboDeFenris{Nordicos{Guerreiro{nome, idade, peso, energia}}}
}

func (l *LoboDeFenris) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var lobosProximos, x, dano int = 0, 1, 40

	for x < len(*lado1) {
		if _, ok := (*lado1)[x].(*LoboDeFenris); ok {
			x++
			lobosProximos++
		} else {
			break
		}
	}

	dano += int(float64(dano)*0.2) * lobosProximos
	(*lado2)[0].ReceberDano(dano)
	return true
}

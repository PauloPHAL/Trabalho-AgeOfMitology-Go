package domain

type Hidra struct {
	Gregos
	Cabecas int
}

func NewHidra(nome string, idade int, peso int, energia int) *Hidra {
	return &Hidra{Gregos{Guerreiro{nome, idade, peso, energia}}, 0}
}

func (h *Hidra) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var dano int = 50
	var vida int = 10

	dano += h.Cabecas * 10
	(*lado2)[0].ReceberDano(dano)
	if (*lado2)[0].GetEnergia() <= 0 {
		h.Cabecas++
		if (*lado1)[0].GetEnergia() < 100 {
			switch guerreiro := (*lado1)[0].(type) {
			case *Hidra:
				// Recupera a vida de uma maneira específica para Grego
				//guerreiro.Gregos.RecuperarVidaGrego(vida)
			case *LeaoDaNemeia:
				//guerreiro.Gregos.RecuperarVidaGrego(vida)
			case *Ciclope:
				//guerreiro.Gregos.RecuperarVidaGrego(vida)
			default:
				guerreiro.RecuperarVida(vida)
			}
		}
	}
	return true
}

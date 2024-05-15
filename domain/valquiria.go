package domain

type Valquiria struct {
	Nordicos
}

func NewValquiria(nome string, idade int, peso int, energia int) *Valquiria {
	return &Valquiria{Nordicos{Guerreiro{nome, idade, peso, energia}}}
}

func (v *Valquiria) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var dano int = 20
	(*lado2)[0].ReceberDano(dano)
	verificarIndices(lado1, 1, 20)
	return true
}

func verificarIndices(fila *[]Lutador, indice, vida int) {
	var ind int = len(*fila) - 1
	if ind > indice {
		(*fila)[indice].RecuperarVida(vida)
	}
}

package domain

type LeaoDaNemeia struct {
	Gregos
}

func NewLeaoDaNemeia(nome string, idade int, peso int, energia int) *LeaoDaNemeia {
	return &LeaoDaNemeia{Gregos{Guerreiro{nome, idade, peso, energia}}}
}

func (l *LeaoDaNemeia) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	gerarDano(lado2, 0, 30)
	gerarDano(lado2, 1, 15)
	gerarDano(lado2, 2, 5)
	return true
}

func gerarDano(fila *[]Lutador, indiceParaAtaque, dano int) {
	tamanho := len(*fila)
	if tamanho > indiceParaAtaque {
		(*fila)[indiceParaAtaque].ReceberDano(dano)
	}
}

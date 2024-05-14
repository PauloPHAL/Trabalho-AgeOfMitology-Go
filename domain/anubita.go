package domain

type Anubita struct {
	Egipcios
}

func NewAnubita(nome string, idade int, peso int, energia int) *Anubita {
	return &Anubita{Egipcios{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Anubita) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var dano int = 15
	indice := len(*lado1)
	indice--
	(*lado1)[0].ReceberDano(dano)
	(*lado1)[indice].ReceberDano(dano)
	return true
}

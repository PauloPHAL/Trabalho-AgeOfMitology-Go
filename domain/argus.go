package domain

type Argus struct {
	Atlantes
}

func NewArgus(nome string, idade int, peso int, energia int) *Argus {
	return &Argus{Atlantes{Guerreiro{nome, idade, peso, energia}}}
}

func (a *Argus) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var dano int = (*guerreiro).GetEnergia()
	(*lado1)[0].ReceberDano(dano)
	return true
}

package domain

type Ciclope struct {
	Gregos
}

func NewCiclope(nome string, idade, peso, energia int) *Ciclope {
	return &Ciclope{Gregos{Guerreiro{Nome: nome, Idade: idade, Peso: peso, Energia: energia}}}
}

func (c *Ciclope) Atacar(guerreiro *Lutador, lado1, lado2 *[]Lutador) bool {
	var dano int = 40
	(*lado2)[0].ReceberDano(dano)
	return false
}

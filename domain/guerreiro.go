package domain

type Lutador interface {
	Atacar(lado1, lado2 *[]Lutador)
	GetNome() string
	GetIdade() int
	GetPeso() int
	GetEnergia() int
}

type Guerreiro struct {
	Nome    string
	Idade   int
	Peso    int
	Energia int
}

func (g *Guerreiro) GetNome() string {
	return g.Nome
}

func (g *Guerreiro) GetIdade() int {
	return g.Idade
}

func (g *Guerreiro) GetPeso() int {
	return g.Peso
}

func (g *Guerreiro) GetEnergia() int {
	return g.Energia
}

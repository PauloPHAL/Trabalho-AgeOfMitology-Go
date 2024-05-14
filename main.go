package main

import (
	"github.com/PauloPHAL/Trabalho-AgeOfMitology-Go/dao"
	"github.com/PauloPHAL/Trabalho-AgeOfMitology-Go/domain"
	"github.com/PauloPHAL/Trabalho-AgeOfMitology-Go/gerencia"
)

func main() {
	var lado1 []domain.Lutador
	var lado2 []domain.Lutador

	dao.PreencherSlices1(&lado1)
	dao.PreencherSlices2(&lado2)

	gerencia.Jogar(&lado1, &lado2)
}

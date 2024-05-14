package main

import (
	"Trabalho1/dao"
	"Trabalho1/domain"
	"fmt"
)

func main() {
	var lado1 []domain.Lutador
	var lado2 []domain.Lutador

	dao.PreencherSlices1(&lado1)
	dao.PreencherSlices2(&lado2)

	for _, lutador1 := range lado1 {
		lutador1.Atacar(&lado1, &lado2)
		fmt.Println(lutador1.GetNome())
	}

	fmt.Println("--------------------------")

	for _, lutador2 := range lado2 {
		lutador2.Atacar(&lado1, &lado2)
		fmt.Println(lutador2.GetNome())
	}

}

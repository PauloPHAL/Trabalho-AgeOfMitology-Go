package dao

import (
	"Trabalho1/domain"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PreencherSlices1(lado1 *[]domain.Lutador) {
	file, err := os.Open("./Lado1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var codigo int
	var nome string
	var idade int
	var peso int
	var lutador domain.Lutador

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.Fields(linha)

		if len(partes) >= 4 {
			codigo, _ = strconv.Atoi(partes[0])
			nome = partes[1]
			idade, _ = strconv.Atoi(partes[2])
			peso, _ = strconv.Atoi(partes[3])
		}

		switch codigo {
		case 1:
			lutador = domain.NewCiclope(nome, idade, peso, 100)
		case 2:
			lutador = domain.NewLeaoDaNemeia(nome, idade, peso, 100)
		case 3:
			lutador = domain.NewHidra(nome, idade, peso, 100)
		case 4:
			lutador = domain.NewValquiria(nome, idade, peso, 100)
		case 5:
			lutador = domain.NewLoboDeFenris(nome, idade, peso, 100)
		}

		if lutador != nil {
			*lado1 = append(*lado1, lutador)
		}
	}

}

func PreencherSlices2(lado2 *[]domain.Lutador) {
	file, err := os.Open("./Lado2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var codigo int
	var nome string
	var idade int
	var peso int
	var lutador domain.Lutador

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.Fields(linha)

		if len(partes) >= 4 {
			codigo, _ = strconv.Atoi(partes[0])
			nome = partes[1]
			idade, _ = strconv.Atoi(partes[2])
			peso, _ = strconv.Atoi(partes[3])
		}

		switch codigo {
		case 1:
			lutador = domain.NewPrometeano(nome, idade, peso, 100)
		case 2:
			lutador = domain.NewSatiro(nome, idade, peso, 100)
		case 3:
			lutador = domain.NewArgus(nome, idade, peso, 100)
		case 4:
			lutador = domain.NewAnubita(nome, idade, peso, 100)
		case 5:
			lutador = domain.NewMumia(nome, idade, peso, 100)
		}

		if lutador != nil {
			*lado2 = append(*lado2, lutador)
		}
	}
}

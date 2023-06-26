package main

import (
	"fmt"
	"math"
)

// Função para calcular a pontuação fuzzy
func calcularPontuacao(sentimento float64) float64 {
	if sentimento < 0.2 {
		return 0.0
	} else if sentimento < 0.5 {
		return (sentimento - 0.2) / 0.3
	} else if sentimento < 0.8 {
		return 1.0
	} else if sentimento <= 1.0 {
		return (1.0 - sentimento) / 0.2
	}
	return 0.0
}

// Função de teste
func testeCalculoPontuacao() {
	testCases := []struct {
		sentimento        float64
		pontuacaoEsperada float64
	}{
		{0.1, 0.0},
		{0.3, 0.3},
		{0.6, 1.0},
		{0.9, 0.5},
		{1.2, 0.0},
	}

	fmt.Println("Teste de Cálculo de Pontuação Fuzzy:")
	for _, testCase := range testCases {
		pontuacao := calcularPontuacao(testCase.sentimento)
		pontuacaoArredondada := math.Round(pontuacao*100) / 100 // Arredondar para duas casas decimais

		if pontuacaoArredondada == testCase.pontuacaoEsperada {
			fmt.Printf("Sentimento: %.2f | Pontuação: %.2f [PASSOU]\n", testCase.sentimento, pontuacaoArredondada)
		} else {
			fmt.Printf("Sentimento: %.2f | Pontuação: %.2f [FALHOU]\n", testCase.sentimento, pontuacaoArredondada)
		}
	}
}

func main() {
	testeCalculoPontuacao()
}

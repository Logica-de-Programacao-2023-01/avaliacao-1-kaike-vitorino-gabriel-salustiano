package q3

import "errors"

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

func DominoPieces(m, n int) (int, error) {
	// Seu código aqui
	if m <= 0 || n <= 0 {
		return 0, errors.New("Ta errado")
	}
	NumSlot := m * n
	Resultado := 0

	if NumSlot%2 == 0 {
		Resultado = NumSlot / 2
		return Resultado, nil
	} else {
		Resultado = NumSlot - 1
		Resultado = Resultado / 2
		return Resultado, nil
	}

	return 0, nil
}

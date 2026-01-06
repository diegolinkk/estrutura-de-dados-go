package main

import (
	"arrays/operacoes"
	"fmt"
)

func main() {
	var numeros [10]int
	for i := 0; i < len(numeros); i++ {
		numeros[i] = i + 1
	}
	fmt.Println("Lista original: ", numeros)
	fmt.Println("EncontrarElemento 15: ", operacoes.EncontrarElemento(numeros[:], 15))
	fmt.Println("EncontrarElemento 52: ", operacoes.EncontrarElemento(numeros[:], 52))
	fmt.Println("AddElementoFinal 52: ", operacoes.AddElementoFinal(numeros[:], 52))
	fmt.Println("AddElementoComeco 431: ", operacoes.AddElementoComeco(numeros[:], 431))
	fmt.Println("RemoverElementoFinal: ", operacoes.RemoverElementoFinal(numeros[:]))
	fmt.Println("RemoverElementoComeco: ", operacoes.RemoverElementoComeco(numeros[:]))
	fmt.Println("AddElIndiceEspecifico: ", operacoes.AddElemIndiceEspecifico(numeros[:], 5, 99))
	fmt.Println("RemElemIndiceEspecifico: ", operacoes.RemElemIndiceEspecifico(numeros[:], 4))
}

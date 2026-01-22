package arrays

import (
	"fmt"
)

func ExemplosArrays() {
	var numeros [10]int
	for i := 0; i < len(numeros); i++ {
		numeros[i] = i + 1
	}
	fmt.Println("Lista original: ", numeros)
	fmt.Println("EncontrarElemento 15: ", EncontrarElemento(numeros[:], 15))
	fmt.Println("EncontrarElemento 52: ", EncontrarElemento(numeros[:], 52))
	fmt.Println("AddElementoFinal 52: ", AddElementoFinal(numeros[:], 52))
	fmt.Println("AddElementoComeco 431: ", AddElementoComeco(numeros[:], 431))
	fmt.Println("RemoverElementoFinal: ", RemoverElementoFinal(numeros[:]))
	fmt.Println("RemoverElementoComeco: ", RemoverElementoComeco(numeros[:]))
	fmt.Println("AddElIndiceEspecifico: ", AddElemIndiceEspecifico(numeros[:], 5, 99))
	fmt.Println("RemElemIndiceEspecifico: ", RemElemIndiceEspecifico(numeros[:], 4))
}

package main

import (
	"data_structures/arrays"
	"data_structures/deque"
	"data_structures/fila"
	"data_structures/pilha"
	"fmt"
)

func main() {
	arrays.ExemplosArrays()
	pilha.ExemplosPilha()
	fmt.Println(pilha.CalculadoraBinario(33))
	fila.ExemplosFila()
	deque.ExemplosDeque()
}

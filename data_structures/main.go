package main

import (
	"data_structures/linkedlist"
	"fmt"
)

func main() {
	// arrays.ExemplosArrays()
	// pilha.ExemplosPilha()
	// fmt.Println(pilha.CalculadoraBinario(33))
	// fila.ExemplosFila()
	// deque.ExemplosDeque()
	lista := linkedlist.Linkedlst[string]{}
	lista.PushBack("Fumiga")
	lista.PushFront("Cachorro")
	lista.PushFront("Gato")
	lista.PushFront("Papagaio")
	fmt.Println(lista.Head.Value)
	lista.PushBack("Zebra")
	lista.PushBack("Texugo")
	lista.PushFront("Cacauga")
	fmt.Println(lista)
	fmt.Println("Removendo item:", lista.PopFront())
	fmt.Println(lista)
	fmt.Println("Procurando Item Inexistente: ", lista.Search("Item Inexistente"))
	fmt.Println("Procurando Fumiga: ", lista.Search("Fumiga"))
	fmt.Println("Procurando Texugo: ", lista.Search("Texugo"))
	fmt.Println("Acessando três nodes adentro (só mostrando que é possível)", lista.Head.Next.Next.Value)
}

package linkedlist

import "fmt"

type Linkedlst[T comparable] struct {
	Head *Node[T]
	Len  int
}

//string para sobrepor o print padrão e imprimir de fato a lista
func (l Linkedlst[T]) String() string {
	if l.Head == nil {
		return "[lista vazia]"
	}

	retorno := "[ "
	currentNode := l.Head
	retorno += fmt.Sprintf("%v", currentNode.Value)

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		retorno += fmt.Sprintf("-> %v", currentNode.Value)
	}

	retorno += " ]"
	return retorno
}

// PushFront	Adiciona um nó no início.	O novo nó vira o Head.
func (l *Linkedlst[T]) PushFront(value T) {
	newNode := Node[T]{
		Value: value,
		Next:  nil,
	}
	if l.Head != nil {
		newNode.Next = l.Head
	}
	l.Len += 1
	l.Head = &newNode
}

// PushBack	Adiciona um nó no fim.	Precisa percorrer até o nil (ou usar o Tail).
func (l *Linkedlst[T]) PushBack(value T) {
	//cria novo node
	newNode := Node[T]{Value: value}
	//se não houver itens, o último na vdd é o primeiro
	if l.Head == nil {
		l.Head = &newNode
		l.Len += 1
		return
	}

	//adiciona node como ultimo item
	lastNode := l.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = &newNode
	l.Len += 1
}

// PopFront	Remove o primeiro item.	O Head pula para o Next.
func (l *Linkedlst[T]) PopFront() Node[T] {
	retorno := *l.Head
	if l.Len <= 1 {
		l.Head = nil
	} else {
		l.Head = l.Head.Next
	}
	return retorno
}

// Search	Procura um valor.	Percorre a lista comparando cada Value.
func (l *Linkedlst[T]) Search(value T) int {
	index := 0
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return index
		}
		index++
		currentNode = currentNode.Next
	}
	return -1
}

package pilha

type Pilha[T any] struct {
	items []T
}

// Push adiciona um novo item ao final da lista
func (p *Pilha[T]) Push(item T) {
	p.items = append(p.items, item)
}

// Pop retorna o último item da lista e o remove da items
func (p *Pilha[T]) Pop() T {
	size := p.Size()
	if size == 0 || p.items == nil {
		var retorno T
		return retorno
	}

	item := p.items[size-1]
	p.items = p.items[:size-1]
	return item
}

// Peek retorna topo da items, mas não remove o elemento
func (p *Pilha[T]) Peek() T {
	size := p.Size()
	if size == 0 || p.items == nil {
		var retorno T
		return retorno
	}
	return p.items[size-1]
}

// isEmpty retorna true se for vazio ou false se tiver algum item
func (p Pilha[T]) IsEmpty() bool {
	return len(p.items) == 0
}

// Clear limpa a lista
func (p *Pilha[T]) Clear() {
	p.items = nil
}

func (p Pilha[T]) Size() int {
	return len(p.items)
}

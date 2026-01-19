package pilha

type Pilha struct {
	items []any
}

// Push adiciona um novo item ao final da lista
func (p *Pilha) Push(item any) {
	p.items = append(p.items, item)
}

// Pop retorna o último item da lista e o remove da items
func (p *Pilha) Pop() any {
	size := p.Size()
	if size == 0 || p.items == nil {
		return nil
	}

	item := p.items[size-1]
	p.items = p.items[:size-1]
	return item
}

// Peek retorna topo da items, mas não remove o elemento
func (p *Pilha) Peek() any {
	size := p.Size()
	if size == 0 || p.items == nil {
		return nil
	}
	return p.items[size-1]
}

// isEmpty retorna true se for vazio ou false se tiver algum item
func (p Pilha) IsEmpty() bool {
	return len(p.items) == 0
}

// Clear limpa a lista
func (p *Pilha) Clear() {
	p.items = nil
}

func (p Pilha) Size() int {
	return len(p.items)
}

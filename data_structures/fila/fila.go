package fila

type Fila[T any] struct {
	itens []T
}

// Size
func (f Fila[T]) Size() int {
	return len(f.itens)
}

// Enqueue
func (f *Fila[T]) Enqueue(item T) {
	f.itens = append(f.itens, item)
}

func (f *Fila[T]) Dequeue() T {
	var retorno T

	if f.IsEmpty() {
		return retorno
	}

	retorno = f.itens[0]
	f.itens = f.itens[1:]
	return retorno
}

// Peek
func (f Fila[T]) Peek() T {
	if f.IsEmpty() {
		var retorno T
		return retorno
	}
	return f.itens[0]
}

// isEmpty
func (f Fila[T]) IsEmpty() bool {
	return len(f.itens) == 0
}

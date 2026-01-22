package deque

type Deque[T any] struct {
	itens []T
}

func (d Deque[T]) IsEmtpy() bool {
	return len(d.itens) == 0
}

// AddFirst
func (d *Deque[T]) AddFirst(item T) {
	// d.itens = append([]T{item}, d.itens...) // abordagem 1 usando spread operator
	//abordagem 2. Expande + desloca o slice
	var zeroValue T
	d.itens = append(d.itens, zeroValue) // expandindo a lista
	copy(d.itens[1:], d.itens[0:])
	d.itens[0] = item
}

// AddLast
func (d *Deque[T]) AddLast(item T) {
	d.itens = append(d.itens, item)
}

// RemoveFirst
func (d *Deque[T]) RemoveFirst() T {
	var retorno T
	if d.IsEmpty() {
		return retorno
	}
	retorno = d.itens[0]
	d.itens = d.itens[1:]
	return retorno
}

// RemoveLast
func (d *Deque[T]) RemoveLast() T {
	var retorno T
	if d.IsEmpty() {
		return retorno
	}
	retorno = d.itens[len(d.itens)-1]
	d.itens = d.itens[:len(d.itens)-1]
	return retorno
}

// PeekFirst
func (d *Deque[T]) PeekFirst() T {
	var retorno T
	if d.IsEmpty() {
		return retorno
	}
	return d.itens[0]
}

// PeekLast
func (d *Deque[T]) PeekLast() T {
	if d.IsEmpty() {
		var retorno T
		return retorno
	}
	return d.itens[len(d.itens)-1]
}

// IsEmtpy
func (d *Deque[T]) IsEmpty() bool {
	return len(d.itens) == 0
}

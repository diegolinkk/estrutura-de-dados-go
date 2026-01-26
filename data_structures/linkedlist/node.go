package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

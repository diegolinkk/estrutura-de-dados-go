package deque

import "fmt"

func ExemplosDeque() {
	fmt.Println("Exemplos Deque: ")
	animais := Deque[string]{}
	animais.AddFirst("Cachorro")
	animais.AddFirst("Gato")
	animais.AddFirst("Papagaio")
	fmt.Println(animais)
	fmt.Println("Peek First: ", animais.PeekFirst())
	fmt.Println("Peek Last: ", animais.PeekLast())
	fmt.Println("Remove First: ", animais.RemoveFirst())
	fmt.Println(animais)
	fmt.Println("Remove Last:", animais.RemoveLast())
	fmt.Println(animais)
}

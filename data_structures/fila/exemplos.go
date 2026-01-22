package fila

import "fmt"

func ExemplosFila() {
	fmt.Println("EXEMPLOS DA ESTRUTURA PILHA: ")
	animais := Fila[string]{}
	animais.Enqueue("Cachorro")
	animais.Enqueue("Gato")
	animais.Enqueue("Papagaio")
	fmt.Println(animais.IsEmpty()) //false
	fmt.Println(animais.Size())    //3
	fmt.Println(animais.Peek())    //cachorro
	fmt.Println(animais.Dequeue()) //cachorro
	fmt.Println(animais.Size())    //2
	fmt.Println(animais.Dequeue()) //gato
	fmt.Println(animais.Dequeue()) //papagaio
	fmt.Println(animais.IsEmpty()) //true
}

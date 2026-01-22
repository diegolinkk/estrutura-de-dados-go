package pilha

import (
	"fmt"
	"strconv"
)

func ExemplosPilha() {
	var animais Pilha[string]

	//adicionando itens
	animais.Push("Cachorro")
	animais.Push("Gato")

	//removendo e olhando itens com Pop
	fmt.Println(animais.Pop()) //Gato
	fmt.Println(animais.Pop()) //Cachorro

	//agora a lista estará vazia, logo isEmtpy será true
	fmt.Println("Is Empty: ", animais.IsEmpty())

	//adicionando os itens de volta
	animais.Push("Cachorro")
	animais.Push("Gato")

	//agora IsEmpty não mais será true
	fmt.Println("Is Empty: ", animais.IsEmpty())

	//Peek olha o topo mas não remove da lista
	fmt.Println(animais.Peek()) //Gato
	fmt.Println(animais.Peek()) //Gato

	//aqui o tamanho continua 2 porque Peek não remove nada
	fmt.Println("Size antes do Clear: ", animais.Size())

	animais.Clear()

	fmt.Println("Is Empty depois do clear: ", animais.IsEmpty()) //true
	fmt.Println("Size depois do Clear: ", animais.Size())        //0

	animais.Push("Orangotango")
	fmt.Println(animais.Peek()) //Orangotango
}

func CalculadoraBinario(numero int) string {
	if numero == 0 {
		return "0"
	}

	resultado := ""
	var p Pilha[int]
	for numero > 0 {
		p.Push(numero % 2)
		numero = numero / 2
	}
	for p.Size() > 0 {
		retorno := p.Pop()
		resultado += strconv.Itoa(retorno)

	}
	return resultado
}

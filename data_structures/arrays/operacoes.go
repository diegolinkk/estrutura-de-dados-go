package arrays

//Recebe o array e um valor para busca e retorna o indice daquele elemento.
//Caso não localizado, retorna -1
func EncontrarElemento(a []int, elemento int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == elemento {
			return i
		}
	}
	return -1
}

//A ABORDAGEM AQUI SERÁ A DE CRIAR UM NOVO ELEMENTO, APESAR DE EXISTIREM OS SLICES
//PARA CONTORNAR A LIMITAÇÃO DO TAMANHO NO RETORNO DE FUNÇÃO, EU VOU TRABALHAR COM ARRAY E RETORNAR UM SLICE DESSE ARRAY

//AddElementoFinal adicionando elemento no final da lista informada e retorna um novo slice
func AddElementoFinal(lista []int, novoElemento int) []int {
	tamanhoLista := len(lista)
	novaLista := make([]int, tamanhoLista+1)
	for i := 0; i < len(lista); i++ {
		novaLista[i] = lista[i]
	}
	novaLista[tamanhoLista] = novoElemento
	return novaLista
}

//adicionando elemento no começo
func AddElementoComeco(lista []int, novoElemento int) []int {
	tamanhoLista := len(lista)
	var novaLista []int = make([]int, tamanhoLista+1)
	novaLista[0] = novoElemento
	for i := 0; i < tamanhoLista; i++ {
		novaLista[i+1] = lista[i]
	}
	return novaLista
}

//removendo elemento do final
func RemoverElementoFinal(lista []int) []int {
	tamanhoLista := len(lista)
	novaLista := make([]int, tamanhoLista-1)
	for i := 0; i < tamanhoLista-1; i++ {
		novaLista[i] = lista[i]
	}
	return novaLista
}

//removendo elemento do começo
func RemoverElementoComeco(lista []int) []int {
	var tamanhoLista int = len(lista)
	var novaLista []int = make([]int, tamanhoLista-1)
	for i := 1; i < tamanhoLista; i++ {
		novaLista[i-1] = lista[i]
	}
	return novaLista
}

//adicionando elemento no indice especifico
func AddElemIndiceEspecifico(lista []int, posicao, elemento int) []int {
	var tamanhoLista int = len(lista)
	novaLista := make([]int, tamanhoLista+1)
	for i := 0; i < posicao; i++ {
		novaLista[i] = lista[i]
	}
	novaLista[posicao] = elemento
	for i := posicao; i < tamanhoLista; i++ {
		novaLista[i+1] = lista[i]
	}
	return novaLista
}

//removendo elemento no indice especifico
func RemElemIndiceEspecifico(lista []int, posicao int) []int {
	tamanhoLista := len(lista)
	novaLista := make([]int, tamanhoLista-1)
	for i := 0; i < posicao; i++ {
		novaLista[i] = lista[i]
	}

	for i := posicao + 1; i < tamanhoLista; i++ {
		novaLista[i-1] = lista[i]
	}
	return novaLista
}

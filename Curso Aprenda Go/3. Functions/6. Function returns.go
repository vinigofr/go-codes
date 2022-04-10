package main

import "fmt"

// Na order
// 1. Primeiro o nome da função pai (retornadora)
// 2. Tipo de retorno (func())
// 3. Informamos também o tipo de retorno que a função returnada retorna. (????)
func returnFunction(x int) func() int {
	y := x * 10

	// Função interna que retorna o inteiro
	return func() int {
		return y
	}
}

func main() {
	// Podemos armazenar a função aqui:
	returnedFunction := returnFunction(20)

	valor := returnedFunction()
	fmt.Println(valor)
}
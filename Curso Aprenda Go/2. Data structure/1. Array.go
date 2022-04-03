package main

import "fmt"

// Declarando um array com um tamanho e tipo específicos
var x [5]int

func main() {
	// Atribuindo valores para cada posição
	x[0] = 5
	x[1] = 5
	x[2] = 5

	// Podemos imprimir a posição específica
	fmt.Println(x[0])

	// Podemos também imprimir a fila inteira
	fmt.Println(x)

	// As posições que não tiverem valores atribuídos começarão em 0
	fmt.Println(x[4])

	// Se verificarmos o tipo de X, será [5]int
	fmt.Printf("%T\n", x)

	// Podemos verificar o tamanho de um array
	println(len(x))

}
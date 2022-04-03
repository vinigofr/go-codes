package main

import "fmt"

func main() {
	// Uma maneira interessante de se obter matrizes é utulizando slices de slices
	// Sua estrutura se parece com isso abaixo.

	multidimensional := [][]string{
		{"oi", "oi"}, // Linha
		{"oi", "oi"}, // Linha
		{"oi", "oi"}, // Linha
// Col. 1 - Col. 2
	}
 
	fmt.Println(multidimensional)
}
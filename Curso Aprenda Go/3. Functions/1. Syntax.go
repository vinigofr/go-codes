package main

import "fmt"

// Sintaxe de uma função simples
func soma(x, y int) int {
	return x + y
}

// Sintaxe de função com multiplos returnos
func multiReturn() (int, string) {
	return 1, ""
}

// Recebento múltiplos parâmetros e retornando multiplos valores
func multiplosParams(multi ...int) (int, int) {
	total := 0
	for _, value := range multi {
		total = total + value
	}
	return total, len(multi)
}

func main() {
	fmt.Println(soma(100,100))
	fmt.Println(multiReturn())
	fmt.Println(multiplosParams(1,2,3,4,5,6,7,8,9,10))
}

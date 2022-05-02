package main

import "fmt"

func main() {
	f()
	fmt.Println("Retornando normalmente de f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado em f", r)
		}
	}()

	fmt.Println("Chamando g")

	g(0)
	fmt.Println("Retornando normalmente de g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Tocando o terror")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer em função g", i)
	fmt.Println("Imprimindo em funcão g", i)
	g(i + 1)
}
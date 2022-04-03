package main

import "fmt"

// O método make cria um slice/array de tamanho específico.
// Ao receber o append quando estiver no seu limite, todos os dados são transmitidos a um novo slice, aumentando um pouco o custo computacional
//
func main() {
  //						type  len capacity
	slice := make([]int, 5, 5)

	// Mostra capacidade
	fmt.Println(cap(slice)) // 5

	slice = append(slice, 1)
	fmt.Println(cap(slice)) // 10
}
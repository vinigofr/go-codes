package main

import "fmt"

func main() {
	// [Início:Fim] Resgata o primeiro elemento a partir do índice informado
	sabores := []string{
		"calabresa",
		"queijo",
		"frango",
		"marguerita",
	}

	fatia := sabores[1:3]
	fmt.Println(fatia)

	for index := 0; index < len(sabores); index ++ {
		fmt.Println(sabores[index]);
	}

	// Elimando valores de acordo com o range, eliminando grango
	// Funciona basicamente selecionando uma parcela que queremos e ignorando a que não queremos
	novo_cardapio := append(sabores[0:2], sabores[3:]...)
	fmt.Printf("%v", novo_cardapio);
}
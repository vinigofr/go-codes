// Muitos paracidos com array, porem mais flexíveis.
// Slide é um tipo de dado composto

package main

import (
	"fmt"
)

func main() {
	// Criando um slice
	slice := []string{"vinicius", "gouveia"}

	// Adicionando itens ao slice com o append
	// O valor é deve ser atribuído a uma nova variável
	slice = append(slice, "teste")
	fmt.Println(slice)

	fruits := []string{"acerola", "tangerina", "goiaba", "manga"}

	for i, valor := range fruits {
		fmt.Println("No índice", i, "há o valor de", valor )
	}
}
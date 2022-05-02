package main

import (
	"errors"
	"fmt"
)

func main() {
	result := errors.New("Erro aqui") // Tipo "error"
	fmt.Println(result)

	anotherResult := fmt.Errorf("Error?")
	fmt.Println(anotherResult)

	// Ou seja, você pode simplesmente returnar algo do tipo error quando quiser
}

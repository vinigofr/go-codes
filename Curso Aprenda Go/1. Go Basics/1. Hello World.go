package main

import "fmt"

func main() {
	// Sempre pacote.instrução
	// No caso abaixo temos um _ (underscore). Isso significa que
	// ali há uma desustruração da função fmt.Println e não iremos
	// utilizar a variável de erro.
	byteNumbers, _ := fmt.Println("Hello World", "Ciao World")

	//  O println retorna um número de bytes e erro
	fmt.Println(byteNumbers)

}
package main

import "fmt"

// Jamais ignore erros com _ (underscore). Pois uma hora vai dar ruim
// Sempre verifique seus erros
func main() {
	// Para funcões que dificilmente retornarão erros, pode-se usar o _
	fmt.Println("Demonstrando Println")
	bytes, _ := fmt.Println("-> 'Value' <-")
	fmt.Printf("%v <- bytes da mensagem\n\n", bytes)


	fmt.Println("Demonstrando Scan")
	var scan1 string
	bytess, err := fmt.Scan(&scan1)
	println(bytess)
	if err != nil {
		panic(err)
	}
}
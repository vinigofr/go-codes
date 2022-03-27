// Desafio

// Utilizando format printing, imprima dos números 33 a 122
// Esses números devem ter convertidos para seus respectivos caracteres.
// Number >> ASCII

package main

import "fmt"

// Para resolver, basta apenas converter o número utilizando o método string()
// https://www.socketloop.com/tutorials/golang-how-to-convert-character-to-ascii-and-back#:~:text=To%20get%20the%20ASCII%20value,integer%20with%20string()%20function

func main() {
	for i := 32; i <= 122; i++ {
		character := string(rune(i))
		fmt.Printf("O número %v em texto é %v\n", i, character);
	}
}
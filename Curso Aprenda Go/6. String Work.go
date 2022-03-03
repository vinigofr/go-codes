package main

import (
	"fmt"
)

// Strings sao imutaveis
// A solucao eh criar uma nova string baseada na anterior
func main() {
	s := "ℋello, Vinicius"
	fmt.Printf("%v\n%T\n", s, s)

	// String convertida em um slice de bytes
	sb := []byte(s)
	// fmt.Printf("%v\n%T", sb, sb)

	// Range entregara caractere por caractere
	for _, value := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", value, value, value, value)
	}

	// Aqui entregara byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i],s[i],s[i],s[i])
	}
	
}
package main

import "fmt"

// Põe na tela todas as letras maiúsculas do alfabeto
// Três vezes por letra.
/* Ex:
65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
*/
func main() {
	for ASCII := 65; ASCII <= 90; ASCII++ {
		fmt.Println(ASCII);

		scape := 1
		for scape <= 3 {
			scape++
			fmt.Printf("\t\t%#U\n", ASCII)
		}
	}
}
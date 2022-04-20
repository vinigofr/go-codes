// Podemos também converter um canal.

package main

import "fmt"

func main() {
	biChannel := make(chan int)

	fmt.Printf("O novo tipo de biChannel é %T\n", (<-chan int)(biChannel))
	fmt.Printf("O novo tipo de biChannel é %T\n", (chan<- int)(biChannel))
}
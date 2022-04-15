package main

import (
	"fmt"
	"sort"
)

func main() {
	// Desordem
	strings := []string{"vinicius", "teste"}
	ints := []int{9,8,7,6,5,4,3,2,1}

	// Ordem
	sort.Strings(strings)
	sort.Ints(ints)

	// Mostra
	fmt.Println(strings)
	fmt.Println(ints)
}
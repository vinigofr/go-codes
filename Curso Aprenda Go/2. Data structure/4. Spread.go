package main

import "fmt"

func main() {
	slice_um := []int{1, 2, 3, 4}
	slice_dois := []int{5, 6, 7}

	// O spread significa espalhar. Então quando quisermos juntar dois slices
	// Podemos utilizar a sintaxe abaixo:
	slice_um = append(slice_um, slice_dois...)

	// Que seria a mesma coisa de fazer isso:
	slice_um = append(slice_um, 5, 6, 7)

	fmt.Print(slice_um)
}
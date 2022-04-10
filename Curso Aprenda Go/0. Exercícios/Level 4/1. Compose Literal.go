package main

import "fmt"

var anArray [5]int

func main() {
	// 1
	anArray[0] = 1
	anArray[1] = 2
	anArray[2] = 3
	anArray[3] = 4
	anArray[4] = 5

	for _, value := range anArray {
		fmt.Println(value)
	}
	fmt.Printf("%T", anArray);
	fmt.Println()

	// 2
	sliced := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range sliced {
		fmt.Println(value)
	}
	fmt.Printf("%T", sliced);

	// Imprima:
	// 1. Do primeiro ao terceiro item de sliced
	newSliced := append(sliced[:3])
	fmt.Println(newSliced)

	// 2. Do quinto ao ultimo
	newSliced = append(sliced[4:])
	fmt.Println(newSliced)

	// 3. Do segundo ao sétimo
	newSliced = append(sliced[1:7])
	fmt.Println(newSliced)

	newSliced = append(sliced[2:len(sliced) - 1])
	fmt.Println(newSliced)
}
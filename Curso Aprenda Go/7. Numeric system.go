// É importante entender um pouco de numeração:
// Binária,
// Hex,
// Dec

package main

import "fmt"

func main() {
	a := 100
	fmt.Printf("%d\t %b\t %#x", a, a, a) // 100 1100100 0x64
}
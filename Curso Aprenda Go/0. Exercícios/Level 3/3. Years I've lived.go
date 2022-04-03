package main

import "fmt"

// Demonstra os anos desde quando você nasceu até agora
func main() {
	bornYear := 2000
	for bornYear <= 2022 {
		fmt.Println(bornYear)
		bornYear++
	}
}
package main

import "fmt"

// Demonstra os anos desde quando você nasceu até agora
func main() {
	bornYear := 2000
	for {
		fmt.Println(bornYear)
		bornYear++
		if !(bornYear <= 2022) {
			break
		}
	}
}
package main

import "fmt"

func main() {
	theMap := map[string]string{
		"vini": "vini",
		"cius": "cius",
	}

	delete(theMap, "cius")

	fmt.Println(theMap)

}
package main

import "fmt"

func main() {
	anonimous := struct {
		aMap   map[string]string
		aSlice []int
	}{
		aMap: map[string]string{"aaa": "aaa"},
		aSlice: []int{1},
	}

	fmt.Println(anonimous)
}
package main

import "fmt"

func returnABool(function func(int) bool) {

}

func aaaa(number int) bool {
	fmt.Println(number)
	return false
}

func main() {
	returnABool(aaaa)
}
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z) // 0 "" e false

	x = 42
	y = "James Bond"
	z = true
}
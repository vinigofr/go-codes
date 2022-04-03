package main

var x interface{}

func main() {
	x = 10
	switch x.(type) {
	case int:
		println("É um inteiro")
	case bool:
		println("É um bool")
	case string:
		println("É um string")
	case float64:
		println("É um float")
	}
}
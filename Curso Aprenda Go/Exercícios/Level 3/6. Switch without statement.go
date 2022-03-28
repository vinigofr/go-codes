package main

func main() {
	switch {
	case 10 == 10:
		println("Deu 10")
		fallthrough
	case 20 == 20:
		println("Deu 20")
	}
}
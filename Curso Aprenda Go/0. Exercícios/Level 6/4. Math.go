package main

import "fmt"

type circulo struct {
	raio float64
}

type info interface {
	area()
}

func (c circulo) area() {
	fmt.Println(c.raio * c.raio * 3.14)
}

func medida(i info) {
	i.area()
}

func main() {

	circle := circulo{
		raio: 2.5,
	}

	circle.area()
	medida(circle)
}
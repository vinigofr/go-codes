package main

import "testing"

func TestSoma(t *testing.T) {
	valores := []int{1,2,3}
	x := Soma(valores...)
	if x != 6 {
		t.Errorf("A soma de %v deveria ser %v", valores, 6)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Soma(1,2,3,4,5,6,7,8,9,10,10,11,4,84,894,984,654,984,65,465468,489,489,498,498,489,49,465,465,4654,654)
	}
}
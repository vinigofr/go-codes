package main

import "testing"

func TestSoma(t *testing.T) {
	valores := []int{1,2,3}
	x := Soma(valores...)
	if x != 6 {
		t.Errorf("A soma de %v deveria ser %v", valores, 6)
	}
}
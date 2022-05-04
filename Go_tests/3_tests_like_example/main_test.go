package main

import (
	"fmt"
	"testing"
)

func ExampleSoma(t *testing.T) {
	fmt.Println(Soma(1,2,3,4,5,7,8,9))
	// Output: 15
}

func TestSoma(t *testing.T) {
	valores := []int{1,2,3}
	x := Soma(valores...)
	if x != 6 {
		t.Errorf("A soma de %v deveria ser %v", valores, 6)
	}
}
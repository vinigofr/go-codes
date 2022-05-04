package main

import "testing"

type Test struct {
	data []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []Test{
		{
			data: []int{1,2,3,4,5,6},
			answer: 21,
		},
		{
			data: []int{1000, 1000, 1000},
			answer: 3000,
		},
		{
			data: []int{1,2,3,4,5,6},
			answer: 21,
		},
	}
	
	for _, value := range tests {
		x := Soma(value.data...)
		if x != value.answer {
			t.Errorf("A soma de %v deveria ser %v", value.data, value.answer)
		}
	};
}
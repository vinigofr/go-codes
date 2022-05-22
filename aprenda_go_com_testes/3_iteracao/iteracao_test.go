package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 50)
	esperado := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 50)
	}
}

func ExampleRepetir() {
	fmt.Println(Repetir("a", 5))
	// Output: aaaaa
}

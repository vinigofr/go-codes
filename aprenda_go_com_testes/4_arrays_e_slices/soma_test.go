package __arrays_e_slices

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}
	t.Run("Testa somar um array de tamanh fixo", func(t *testing.T) {
		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}

	})
}

func TestSomaTudo(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8}
	slice3 := []int{9, 10, 11, 12}

	resultado := SomaTudo(slice1, slice2, slice3)
	esperado := []int{10, 26, 42}

	/*
		Podemos verificar os slices um a um através de iteração, mas para este caso
		vamos utilizar o reflect.DeepEqual
		OBS: Ele não causa erro de compilação, por isso, cuidado.
	*/
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v esperado %v", resultado, esperado)
	}
}

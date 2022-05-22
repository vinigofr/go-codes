package __arrays_e_slices

import "testing"

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

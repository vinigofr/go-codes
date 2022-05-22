package __arrays_e_slices

func Soma(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	for _, currentSlice := range numerosParaSomar {
		currentSum := 0
		for _, number := range currentSlice {
			currentSum += number
		}
		somas = append(somas, currentSum)
	}
	return
}

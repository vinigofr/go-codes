package __arrays_e_slices

func Soma(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

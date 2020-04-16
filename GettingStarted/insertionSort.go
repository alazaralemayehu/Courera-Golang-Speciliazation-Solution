package GettingStarted

func InsertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		for j := i; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}
	return numbers
}

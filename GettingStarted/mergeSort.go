package GettingStarted

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	} else {
		middle := len(slice) / 2

		left := slice[:middle]
		right := slice[middle:]

		leftResult := MergeSort(left)
		rightResult := MergeSort(right)
		return Merge(leftResult, rightResult)
	}
}
func Merge(left, right []int) []int {
	i, j, size := 0, 0, len(left)+len(right)
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if i <= len(left)-1 && j > len(right)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

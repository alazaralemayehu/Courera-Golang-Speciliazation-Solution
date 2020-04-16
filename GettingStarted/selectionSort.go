package GettingStarted

import "fmt"

func SelectionSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			slice[minIndex], slice[i] = slice[i], slice[minIndex]

		}
	}
	fmt.Println(slice)

}

package GettingStarted

import "fmt"

func BubbleSorts(numbersToBeSorted []int) {

	for i := 0; i < len(numbersToBeSorted); i++ {
		fmt.Println(9)
		for j := 0; j < len(numbersToBeSorted)-i-1; j++ {

			if numbersToBeSorted[j] > numbersToBeSorted[j+1] {
				swap(numbersToBeSorted, j)
			}
		}
	}
	fmt.Println(numbersToBeSorted)
}
func swap(listOfNumbers []int, index int) {
	listOfNumbers[index], listOfNumbers[index+1] = listOfNumbers[index+1], listOfNumbers[index]
}

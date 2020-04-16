package Concurrency

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

//Divides the array into four parts and

func main1() {
	var wg sync.WaitGroup
	var slice = []int{}
	var userInput string

	for {
		fmt.Println("Enter the number, X to stop entering numbers ")
		fmt.Scan(&userInput)
		if userInput == "X" {
			break
		}

		var integerInput, _ = strconv.Atoi(userInput)
		slice = append(slice, integerInput)

	}

	size := len(slice)
	var slice1, slice2, slice3, slice4 []int
	wg.Add(1)
	go BubbleSort(slice[:size/4], &wg, &slice1)
	wg.Add(1)
	go BubbleSort(slice[size/4:size/2], &wg, &slice2)
	wg.Add(1)
	go BubbleSort(slice[size/2:3*size/4], &wg, &slice3)
	wg.Add(1)
	go BubbleSort(slice[3*size/4:], &wg, &slice4)

	wg.Wait()

	finalResult := []int{}
	finalResult = append(finalResult, slice1...)
	finalResult = append(finalResult, slice2...)
	finalResult = append(finalResult, slice3...)
	finalResult = append(finalResult, slice4...)
	sort.Ints(finalResult)
	fmt.Println("The final Merged and Sorted Array is :", finalResult)
}

func BubbleSort(numbersToBeSorted []int, wg *sync.WaitGroup, finalAnswer *[]int) {
	fmt.Println("Sorting ", numbersToBeSorted)
	for i := 0; i < len(numbersToBeSorted); i++ {
		for j := 0; j < len(numbersToBeSorted)-i-1; j++ {

			if numbersToBeSorted[j] > numbersToBeSorted[j+1] {
				swap(numbersToBeSorted, j)
			}
		}
	}
	*finalAnswer = numbersToBeSorted
	fmt.Println("Sorting is Done: The final sorted array is ", *finalAnswer)
	defer wg.Done()
}
func swap(listOfNumbers []int, index int) {
	listOfNumbers[index], listOfNumbers[index+1] = listOfNumbers[index+1], listOfNumbers[index]
}

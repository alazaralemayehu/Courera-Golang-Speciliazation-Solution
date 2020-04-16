package GettingStarted

import (
	"fmt"
	"sort"
	"strconv"
)

func mainn() {
	var l = []int{}
	var userInput string
	for {
		fmt.Println("Enter the number ")
		fmt.Scan(&userInput)
		if userInput == "X" {
			break
		}

		var integerInput, _ = strconv.Atoi(userInput)
		l = append(l, integerInput)
		sort.Ints(l)
		for _, value := range l {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("The final sorted slice is ")
	for _, value := range l {
		fmt.Print(value)
		fmt.Print(" ")
	}
}

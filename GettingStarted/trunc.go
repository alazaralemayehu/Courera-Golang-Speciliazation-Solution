package GettingStarted

import (
	"fmt"
	"strconv"
)

func mai2n() {
	var userInput string
	fmt.Println("please enter any number ")
	fmt.Scan(&userInput)
	fmt.Println(Truncate(userInput))
}
func Truncate(userInput string) int {

	var finalString string = ""
	for _, value := range userInput {
		if string(value) == "." {
			break
		}
		finalString += string(value)
	}
	var finalAnswer, _ = (strconv.Atoi(finalString))
	return (finalAnswer)
}

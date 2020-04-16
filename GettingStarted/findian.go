// This program is writte by Alazar Alemayehu
package GettingStarted

import (
	"fmt"
	"strings"
)

func main1() {
	var userInput string
	fmt.Println("please enter any number ")
	fmt.Scan(&userInput)
	userInput = strings.ToLower(userInput)
	if strings.Contains(userInput, "a") && strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found")
	}
}

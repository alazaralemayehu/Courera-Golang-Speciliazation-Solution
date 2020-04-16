package main

import (
	"fmt"
	"math"
	"strings"
)

func maqin() {

	// s := []int{4, 3, 2, 10, 12, 1, 5, 6}
	// SelectionSort(s)

	// fmt.Println(s)

	// fmt.Println(applyIt(incFn, 10))
	// fmt.Println(applyIt(decFn, 10))

	// Dist1 := MakeDistOrgin(0, 0)
	// Dist2 := MakeDistOrgin(2, 2)

	// fmt.Println(Dist1(2, 2))
	// fmt.Println(Dist2(2, 2))
	var userInput string
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worm",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	// var animal Animal
	for {
		fmt.Println("Pleas enter the name of the animal and its property")
		fmt.Scan(&userInput)
		userRequests := strings.Split(userInput, " ")
		fmt.Println(userRequests)
	}
}

func incFn(x int) int {
	return x + 1
}
func decFn(x int) int {
	return x - 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func MakeDistOrgin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

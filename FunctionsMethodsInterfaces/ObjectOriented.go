package FunctionsMethodsInterfaces

import (
	"fmt"
	"strings"
	"time"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a snimal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}
func (a animal) Speak() {
	fmt.Println(a.noise)
}

func maaain() {

	var animalType string
	var animalProperty string
	var animal Animal

	// var animal Animal
	for {
		fmt.Println("Pleas enter the name of the animal and its property")
		// userRequests := strings.Split(userInput, " ")
		fmt.Scanf("%v %v", &animalType, &animalProperty)
		animalType = strings.ToLower(animalType)
		animalProperty = strings.ToLower(animalProperty)

		var isValidAnimal = true

		if animalType == "cow" {

			animal = Animal{
				food:       "grass",
				locomotion: "walk",
				noise:      "moo",
			}

		} else if animalType == "bird" {
			animal = Animal{
				food:       "worm",
				locomotion: "fly",
				noise:      "peep",
			}
		} else if animalType == "snake" {
			animal = Animal{
				food:       "mice",
				locomotion: "slither",
				noise:      "hsss",
			}
		} else {
			isValidAnimal = false
			fmt.Println("Please enter valid animal name")
		}

		// fmt.Println(animal)

		if isValidAnimal {
			if animalProperty == "eat" {
				animal.Eat()
			} else if animalProperty == "move" {
				animal.Move()
			} else if animalProperty == "speak" {
				animal.Speak()
			} else {
				isValidAnimal = false
				// fmt.Println("Please enter valid animal Property")
			}
		}
		time.Sleep(time.Second * 3)
	}

}

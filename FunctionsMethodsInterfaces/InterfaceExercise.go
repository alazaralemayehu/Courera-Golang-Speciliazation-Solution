package FunctionsMethodsInterfaces

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

type Cow struct {
	name string
}

type Snake struct {
	name string
}

type Bird struct {
	name string
}

func (c Cow) getName() string {
	return c.name
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Cow) Move() {
	fmt.Println("Walk")
}

func (s Snake) getName() string {
	return s.name
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("Slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func (b Bird) getName() string {
	return b.name
}

func (b Bird) Eat() {
	fmt.Println("Worms")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (f Bird) Move() {
	fmt.Println("fly")
}

type Shape interface {
	Area()
	// Permieter() float64
}

var animals []Animal

func mai2n() {
	animals = make([]Animal, 0, 1)
	for {

		fmt.Print("> ")
		var query1 string
		var query2 string
		var query3 string

		fmt.Scanf("%v %v %v", &query1, &query2, &query3)

		query1 = strings.ToLower(query1)
		query2 = strings.ToLower(query2)
		query3 = strings.ToLower(query3)

		if query1 == "newanimal" {
			animalCreated := createNewAnimal(query2, query3)
			if animalCreated {
				fmt.Println("Created it")
			} else {
				fmt.Println("Something is wrong with your input")
			}
		} else if query1 == "query" {
			queryProcessed := processQuery(query2, query3)
			if !queryProcessed {
				fmt.Println("Something is wrong with your input")
			}
		} else {
			fmt.Println("Please input the correct input")
		}
	}
}

func createNewAnimal(animalName, animalType string) bool {
	if animalType == "cow" {
		c := Cow{animalName}
		animals = append(animals, c)
		return true
	} else if animalType == "bird" {
		b := Bird{animalName}
		animals = append(animals, b)
		return true
	} else if animalType == "snake" {
		s := Snake{animalName}
		animals = append(animals, s)
		return true
	}

	return false
}

func processQuery(animalName, queryCharacterstics string) bool {
	var targetAnimal Animal
	for _, animal := range animals {
		if animal.getName() == animalName {
			targetAnimal = animal
			break
		}
	}
	if targetAnimal == nil {
		return false
	}

	if queryCharacterstics == "move" {
		targetAnimal.Move()
	} else if queryCharacterstics == "eat" {
		targetAnimal.Eat()
	} else if queryCharacterstics == "speak" {
		targetAnimal.Speak()
	} else {
		return false
	}
	return true

}

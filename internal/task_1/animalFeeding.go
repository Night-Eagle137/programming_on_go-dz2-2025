package task1

import (
	"bufio"
	"io"
	"strings"
)

type Animal interface { // interface for ours structures
	GetName() string
	WhatDidYouEat() string
	Eat(str rune)
}

type BaseAnimal struct { // create a base structure
	name  string
	diary string // note for animal's meals
}

func (animal *BaseAnimal) GetName() string { // base function to return animal's name
	return animal.name
}

func (animal *BaseAnimal) WhatDidYouEat() string { // base function to return the diary of meals
	return animal.diary
}

type Cat struct {
	BaseAnimal
}

func (cat *Cat) Eat(str rune) { // function for Cat's meals
	if str >= '0' && str <= '9' {
		cat.diary += string(str)
	}
}

type Dog struct {
	BaseAnimal
}

func (dog *Dog) Eat(str rune) { // function for Dog's meals
	if (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') {
		dog.diary += string(str)
	}
}

type Bird struct {
	BaseAnimal
}

func (bird *Bird) Eat(str rune) { // function for Bird's meals
	if !(str >= '0' && str <= '9') && !(str >= 'a' && str <= 'z') && !(str >= 'A' && str <= 'Z') {
		bird.diary += string(str)
	}
}

func CreateAnimals(animalType, animalName string) Animal {
	switch animalType {
	case "cat":
		return &Cat{BaseAnimal{name: animalName}}
	case "dog":
		return &Dog{BaseAnimal{name: animalName}}
	case "bird":
		return &Bird{BaseAnimal{name: animalName}}
	default:
		return nil
	}
}

func AnimalFeeding(stream io.Reader) []string {
	scanner := bufio.NewScanner(stream) // pulling data
	var food string

	if scanner.Scan() {
		food = scanner.Text() // take a feeder with the rune
	}

	ourAnimals := []Animal{} // slice of input animals with their names

	for scanner.Scan() {
		line := scanner.Text() // take new line untill end

		if line == "end" { // cheack for end
			break
		}

		if line == "" { // if line is empty
			continue
		}
		
		lineParts := strings.Fields(line) // split the line

		if len(lineParts) < 2 {
			continue
		}

		animalType := strings.ToLower(lineParts[0]) // unifying the strings for our structure
		animalName := lineParts[1]

		animal := CreateAnimals(animalType, animalName) // getting adress og animal

		if animal != nil {
			ourAnimals = append(ourAnimals, animal) // fill the slice with adress
		}
	}

	foodRunes := []rune(food)
	animalQueue := make([]Animal, len(ourAnimals)) // slice for animal's stack
	copy(animalQueue, ourAnimals)

	// feeding alghorithm

	for len(foodRunes) > 0 {
		currentAnimal := animalQueue[0] // take first animal
		currentSymbol := foodRunes[0] // take first char		
		currentAnimal.Eat(currentSymbol) // eating =)
		foodRunes = foodRunes[1:] // deleting currentSymbol
		animalQueue = append(animalQueue[1:], currentAnimal) // add to the end of slice currentAnimal
	}

	results := []string{}

	for _, animal := range ourAnimals {
		results = append(results, animal.GetName()+" "+animal.WhatDidYouEat())
	}

	return results
}

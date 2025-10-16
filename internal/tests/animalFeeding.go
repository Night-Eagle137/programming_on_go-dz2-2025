package task1

import (
	"bufio"
	"io"
	"strings"
)

type Animal interface { // interface for ours structures
	GetName() string
	WhatDidYouEat() string
	Eat(str rune) bool
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

func (cat *Cat) Eat(str rune) bool { // function for Cat's meals
	if str >= '0' && str <= '9' {
		cat.diary += string(str)
		return true
	}
	return false
}

type Dog struct {
	BaseAnimal
}

func (dog *Dog) Eat(str rune) bool { // function for Dog's meals
	if (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'B') {
		dog.diary += string(str)
		return true
	}
	return false
}

type Bird struct {
	BaseAnimal
}

func (bird *Bird) Eat(str rune) bool { // function for Bird's meals
	if !(str >= '0' && str <= '9') && !(str >= 'a' && str <= 'z') && !(str >= 'A' && str <= 'B') {
		bird.diary += string(str)
		return true
	}
	return false
}

func AnimalFeeding(stream io.Reader) []string {
	scanner := bufio.NewScanner(stream) // pulling data

	food := scanner.Text() // take a feeder

	ourAnimals := make(map[string]string) // map of input animals with their names

	for scanner.Scan() {
		line := scanner.Text() // take new line untill end

		if line == "end" { // cheack for end
			break
		}

		lineParts := strings.Fields(line) // split the line

		animalType := strings.ToLower(lineParts[0]) // unifying the strings for our structure
		animalName := lineParts[1]

		ourAnimals[animalType] = animalName // fill the map
	}

}

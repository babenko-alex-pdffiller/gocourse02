package main

import (
	"fmt"
	"math/rand/v2"
)

type ZooKeeper struct {
	Name string
}

func (keper *ZooKeeper) PutToCage(animal *Animal, cage *Cage) {
	cage.Animal = animal
	animal.IsFree = false
}

type Animal struct {
	Name   AnimalType
	IsFree bool
}

func (a *Animal) runAway() {
	a.IsFree = true
}

type AnimalType string

const (
	Lion   AnimalType = "Lion"
	Tiger  AnimalType = "Tiger"
	Duck   AnimalType = "Duck"
	Monkey AnimalType = "Monkey"
)

type Cage struct {
	Animal *Animal
}

type Zoo struct {
	City  string
	Cages []Cage
}

func NewZoo(city string) *Zoo {
	return &Zoo{City: city}
}

func (cage *Cage) open() {
	cage.Animal = nil
}

func main() {
	animalNames := [4]AnimalType{Lion, Tiger, Duck, Monkey}
	animalsAmount := rand.IntN(10)

	var animals []Animal

	zoo := NewZoo("Kyiv")
	fmt.Printf("There are %d animals in the %s zoo\n", animalsAmount+1, zoo.City)

	zooKeper := ZooKeeper{"Bob"}

	// Create animals and cages
	for i := 0; i <= animalsAmount; i++ {
		name := animalNames[rand.IntN(4)]
		animal := Animal{name, false}
		animals = append(animals, animal)
		fmt.Printf("%s in cage\n", animals[i].Name)
		zoo.Cages = append(zoo.Cages, Cage{&animal})
	}

	// Opening cells randomly
	for i, _ := range zoo.Cages {
		isOpen := rand.IntN(2) == 0
		if isOpen {
			zoo.Cages[i].open()
			animals[i].runAway()
			fmt.Printf("Cage #%d is open, %s is free\n", i, animals[i].Name)
		}
	}

	// ZooKeeper sent free animals to cages
	for i, cage := range zoo.Cages {
		if cage.Animal == nil {
			fmt.Printf("Cage #%d is empty\n", i)
			for a := 0; a <= animalsAmount; a++ {
				if animals[a].IsFree == true {
					fmt.Printf("Keper sent %s to cage\n", animals[a].Name)
					zooKeper.PutToCage(&animals[a], &zoo.Cages[i])
					break
				}
			}

		}
	}
}

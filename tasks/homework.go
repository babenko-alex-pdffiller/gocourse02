package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ZooKeeper struct {
	Name string
}

func (keper *ZooKeeper) SendToCage(animal *Animal, cage *Cage) {
	cage.Animal = animal
	animal.IsFree = false
}

type Animal struct {
	Name   string
	IsFree bool
}

func (a *Animal) makeFree() {
	a.IsFree = true
}

const (
	Lion   = "Lion"
	Tiger  = "Tiger"
	Duck   = "Duck"
	Monkey = "Monkey"
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

func (cage *Cage) openCage() {
	cage.Animal = nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	animalsNames := [4]string{Lion, Tiger, Duck, Monkey}
	animalsAmount := rand.Intn(10)

	var animals []Animal

	zoo := NewZoo("Kyiv")
	fmt.Printf("There are %d animals in the %s zoo\n", animalsAmount+1, zoo.City)

	zooKeper := ZooKeeper{"Bob"}

	// Create animals and cages
	for i := 0; i <= animalsAmount; i++ {
		name := animalsNames[rand.Intn(4)]
		animal := Animal{name, false}
		animals = append(animals, animal)
		fmt.Printf("%s in cage\n", animals[i].Name)
		zoo.Cages = append(zoo.Cages, Cage{&animal})
	}

	// Opening cells randomly
	for i, _ := range zoo.Cages {
		isOpen := rand.Intn(2) == 0
		if isOpen == true {
			zoo.Cages[i].Animal = nil
			animals[i].IsFree = true
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
					zooKeper.SendToCage(&animals[a], &zoo.Cages[i])
					break
				}
			}

		}
	}
}

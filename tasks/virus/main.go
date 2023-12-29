package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Virus - структура, яка представляє вірус
type Virus struct {
	Genome        string
	InfectedCells int
}

// InfectionInterface - інтерфейс для представлення властивостей віруса
type Cell interface {
	BecomeInfected()
	IsInfected() bool
}

// Cell - структура, яка представляє клітину, яку вірус може заражати
type BloodCell struct {
	ID       int
	Infected bool
}

func (c *BloodCell) IsInfected() bool {
	return c.Infected
}

func (c *BloodCell) BecomeInfected() {
	c.Infected = true
}

// NewVirus - функція для створення нового вірусу з заданим геномом
func NewVirus(genome string) *Virus {
	return &Virus{
		Genome: genome,
	}
}

// Infect - метод для зараження клітини вірусом
func (v *Virus) Infect(cell Cell) {
	if !cell.IsInfected() {
		cell.BecomeInfected()
		v.InfectedCells++
	}
}

// Replicate - метод для реплікації вірусу
func (v *Virus) Replicate() *Virus {
	newVirus := &Virus{
		Genome: v.Genome,
	}
	fmt.Println("Virus is replicating...")
	return newVirus
}

// Replicate - метод для реплікації вірусу
func (v *Virus) ReplicateWithMutation(newGenome string) *Virus {
	v.mutate(newGenome)
	newVirus := &Virus{
		Genome: v.Genome,
	}
	fmt.Println("Virus is replicating...")
	return newVirus
}

// Mutate - метод для мутації вірусу
func (v *Virus) mutate(newGenome string) {
	v.Genome = newGenome

	fmt.Println("Virus is mutating...")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	virus := NewVirus("AGCTAGCT")

	// Створення клітин і вірус заражає їх
	for i := 0; i < 10; i++ {
		virus.Infect(&BloodCell{ID: i + 1})
	}

	// Вірус реплікується та мутує
	_ = virus.Replicate()
	_ = virus.ReplicateWithMutation("AGCAAGCT")

	fmt.Printf("Number of infected cells: %d\n", virus.InfectedCells)
}

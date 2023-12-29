package main

import (
	"encoding/json"
	"fmt"
)

var greyColor = `grey`
var blueColor = `blue`

const (
	Room PropertyType = iota
	Bed
	Window
	Table
)

type PropertyType int

func (pt PropertyType) Name() string {
	switch pt {
	case Room:
		return "Room"
	case Bed:
		return "Bed"
	case Window:
		return "Window"
	case Table:
		return "Table"
	default:
		return "Unknown"
	}
}

type HospitalProperty struct {
	Name       string
	Type       PropertyType
	color      string
	additional *HospitalProperty
}

type Accounting struct {
	Prop001 HospitalProperty `json:"prop1"`
	Prop002 HospitalProperty `json:"prop2"`
	Prop003 HospitalProperty `json:"prop3"`
}

func main() {
	// Створення екземпляра структури Accounting із zero value
	var acc1 Accounting

	// Демонстрація серіалізації в JSON та вивід результатів (zero value)
	printMarshalled(acc1)

	// Створення екземплярів властивостей з використанням констант типів
	room := HospitalProperty{
		Name:  Room.Name(),
		Type:  Room,
		color: greyColor,
		additional: &HospitalProperty{
			Name: Table.Name(),
			Type: Table,
		},
	}
	bed := HospitalProperty{Bed.Name(), Bed, blueColor, nil}
	window := HospitalProperty{
		Name:  Window.Name(),
		Type:  Window,
		color: greyColor,
	}

	// Створення екземпляра структури Accounting з використанням створених властивостей
	acc1 = Accounting{
		Prop001: room,
		Prop002: bed,
		Prop003: window,
	}

	// Демонстрація серіалізації в JSON та вивід результатів
	printMarshalled(acc1)
}

func printMarshalled(a Accounting) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		//Помилка серіалізації в JSON:
		return
	}

	//JSON представлення Accounting
	fmt.Println(string(jsonData))
}

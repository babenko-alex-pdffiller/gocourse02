package main

import (
	"fmt"
)

type Patient struct {
	Name string
	Id   int
	Age  float64
}

type File struct {
	patient Patient
	pages   int
}

func (b File) Pages() int {
	return b.pages
}

func (b File) Patient() Patient {
	return b.patient
}

func (b *File) SetPages(pages int) {
	b.pages = pages
}

func (b File) DoNotSetPages(pages int) {
	b.pages = pages
}

type DentistPatientFile struct {
	File
	CariesNumber int
}

func (b DentistPatientFile) FullInfo() string {
	return fmt.Sprintf("%s with id %d has file with %d pages\n", b.patient.Name, b.patient.Id, b.Pages())
}

func (b *DentistPatientFile) SetPages(pages int) {
	fmt.Printf("Child method has been called\n")
	b.File.SetPages(pages)
}

type SurgeonPatientFile struct {
	File
	SurgeriesCount int
}

func (s SurgeonPatientFile) FullInfo() string {
	return fmt.Sprintf("%s with id %d has file with %d pages and had %d surgeries\n", s.patient.Name, s.patient.Id, s.Pages(), s.SurgeriesCount)
}

func (s *SurgeonPatientFile) SetPages(pages int) {
	fmt.Printf("SurgeonPatientFile method has been called\n")
	s.File.SetPages(pages)
}

func main() {
	// Створення екземпляра Patient
	patient := Patient{
		Name: "John Doe",
		Id:   123,
		Age:  30,
	}

	// Створення екземпляра File
	file := File{
		patient: patient,
		pages:   5,
	}

	// Демонстрація методів File
	fmt.Printf("File pages: %d\n", file.Pages())
	fmt.Printf("File patient: %v\n", file.Patient())

	// Спроба виклику методу, який не змінює стан структури
	file.DoNotSetPages(10)
	fmt.Printf("After calling DoNotSetPages: File pages: %d\n", file.Pages())

	// Спроба виклику методу, який змінює стан структури
	file.SetPages(8)
	fmt.Printf("After calling SetPages: File pages: %d\n", file.Pages())

	// Створення екземпляра DentistPatientFile
	dentistPatientFile := DentistPatientFile{
		File:         file,
		CariesNumber: 3,
	}

	// Демонстрація методів DentistPatientFile та перевизначення SetPages
	fmt.Printf("DentistPatientFile full info: %s", dentistPatientFile.FullInfo())
	dentistPatientFile.SetPages(12)

	// Вивід інформації після виклику SetPages
	fmt.Printf("After calling SetPages on DentistPatientFile: File pages: %d\n", dentistPatientFile.Pages())

	surgeonPatientFile := SurgeonPatientFile{
		File:           file,
		SurgeriesCount: 2,
	}

	// Виклик методів FullInfo поліморфно
	printFullInfo(&dentistPatientFile)
	printFullInfo(&surgeonPatientFile)

	// exercise
	// ????
}

// Функція для демонстрації поліморфізму
func printFullInfo(fp FullInfoPrinter) {
	fmt.Println(fp.FullInfo())
}

// Інтерфейс для поліморфізму
type FullInfoPrinter interface {
	FullInfo() string
	SetPages(pages int)
	Pages() int
}

package main

import (
	"fmt"
	"time"
)

type Patient struct {
	name *string
	Id   int
	Age  int
}

func (p *Patient) DeepCopy1() *Patient {
	return &Patient{
		name: p.name,
		Age:  p.Age,
	}
}

func (p *Patient) DeepCopy2() *Patient {
	tmp := *p
	return &tmp
}

// Name returns John Doe if unknown
func (p *Patient) Name() string {
	if p.name == nil {
		return "John Doe"
	} else {
		return *p.name
	}
}

func (p *Patient) Print() {
	fmt.Printf("ID: %d\n", p.Id)
	fmt.Printf("Patient Name: %s\n", p.Name())
	fmt.Printf("Patient Age: %d\n", p.Age)
}

func NewBlankFile(t string) File {
	return File{title: t}
}

type File struct {
	title    string
	patient  *Patient
	pages    int
	birthDay *time.Time
}

func (b File) Pages() int {
	return b.pages
}

func (b File) Patient() *Patient {
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
	return fmt.Sprintf("%s with id %d has file with %d pages\n", b.patient.Name(), b.patient.Id, b.Pages())
}

func (b *DentistPatientFile) SetPages(pages int) {
	fmt.Printf("Child method has been called\n")
	b.File.SetPages(pages)
}

type SurgeonPatientFile struct {
	*File
	SurgeriesCount int
}

func (s SurgeonPatientFile) FullInfo() string {
	return fmt.Sprintf("%s with id %d has file with %d pages and had %d surgeries\n", s.patient.Name(), s.patient.Id, s.Pages(), s.SurgeriesCount)
}

func (s *SurgeonPatientFile) SetPages(pages int) {
	fmt.Printf("SurgeonPatientFile method has been called\n")
	s.File.SetPages(pages)
}

func main() {
	// Створення екземпляра Patient
	pName := "Ernest H"
	patient := Patient{
		name: &pName,
		Id:   123,
		Age:  30,
	}

	// копіювання структури
	//copyForPatient := patient.DeepCopy1()
	copyForPatient := patient.DeepCopy2()
	copyForPatient.Id = 124
	copyForPatient.Age = 33
	patient.Print()
	copyForPatient.Print()
	fmt.Println()

	// Створення екземпляра File
	file := NewBlankFile(`caries`)
	file = File{
		patient: &patient,
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
		File:           &file,
		SurgeriesCount: 2,
	}

	// Виклик методів FullInfo поліморфно
	printFullInfo(&dentistPatientFile)
	printFullInfo(&surgeonPatientFile)
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

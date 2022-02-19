package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

type Dog struct {
	Name  string
	Breed string
}

func main() {
	dog := Dog{
		Name:  "Cheeper",
		Breed: "German Sheperd",
	}

	PrintInfo(dog)

	gorrila := Gorilla{
		Name: "King Kong",
		Color: "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorrila)

}
//declared as parameter d is type of Dog
func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

//declared as parameter g as type of Gorilla
func (g Gorilla) Says() string {
	return "KUKUKAKA"
}

func (g Gorilla) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This Animal Says", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}
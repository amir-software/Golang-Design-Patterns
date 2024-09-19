package main

import (
	"fmt"
)

// Parent Entity (Abstract model)
type Person struct {
	name   string
	gender string
}

func (person Person) addName() {
	fmt.Print("Name Added")
}

func (person Person) addLastName() {
	fmt.Print("Last Name added Added")
}

// ####################
// Child Entities (Concrete model )
type Male struct {
	Person // All the method of Person would be implement automatically
}

type FeMale struct {
	Person
}

// ###########################
// Operation InterFaces (it has the contorl of the diffirent behaviors)
type personInterFace interface {
	addName()
	addLastName()
}

func createMale() personInterFace {
	return Male{
		Person: Person{
			name:   "Amir",
			gender: "Male",
		},
	}
}
func createFeMale() personInterFace {
	return FeMale{
		Person: Person{
			name:   "Hasti",
			gender: "FeMale",
		},
	}
}

// ###########################
// Factory Method
func getPerson(kind string) personInterFace {
	if kind == "MALE" {
		return createMale()
	}
	if kind == "FEMALE" {
		return createFeMale()
	}
	return nil
}

func main() {
	person := getPerson("MALE") // Call a factory
	person.addName()

}

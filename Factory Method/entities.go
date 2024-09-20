package main

import (
	"fmt"
)

// Parent Entity (ABSTRACT model)
type Person struct {
	name   string
	gender string
}

func (person Person) GetName() {
	fmt.Printf("Name :> %s", person.name)
}

// Child Entities (CONCRETE model)
type Male struct {
	Person // All the method of Person would be implement automatically
}

type FeMale struct {
	Person
}

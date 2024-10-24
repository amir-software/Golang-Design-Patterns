package main

import "fmt"

type educationalSection struct {
	next Handler
}

func (edhandler *educationalSection) setNext(next Handler) {
	edhandler.next = next
}

func (edhandler *educationalSection) execute(s *Student) {
	if s.isPassedLibrarySection {
		s.isPassedEducationalSection = true
	} else {
		fmt.Println("The Student Has'n finished his library section")
		return
	}

	fmt.Println("The Student have done it's Education")
	fmt.Print("He is graduating")
}

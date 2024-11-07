package main

import "fmt"

type SecondLine struct {
	med *Mediator
}

func (s *SecondLine) CanGo() bool {
	if s.med.isAllow {
		fmt.Println("The Second line is free to go")
	} else {
		fmt.Println("The Second line is Blocked")
	}

	return s.med.isAllow
}

func (s *SecondLine) Full() {
	s.med.isAllow = false

	fmt.Println("The Second line is full of car")
}

func (s *SecondLine) Free() {
	s.med.isAllow = true

	fmt.Println("The Second line is now free")
}

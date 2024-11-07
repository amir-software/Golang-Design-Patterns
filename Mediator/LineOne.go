package main

import "fmt"

type IStreet interface {
	CanGo()
	Full()
	Free()
}

type FirstLine struct {
	med *Mediator
}

func (f *FirstLine) CanGo() bool {
	if f.med.isAllow {
		fmt.Println("The first line is free to go")

	} else {
		fmt.Println("The first line is Blocked")
	}

	return f.med.isAllow
}

func (f *FirstLine) Full() {
	f.med.isAllow = false

	fmt.Println("The First line is full of car")
}

func (f *FirstLine) Free() {
	f.med.isAllow = true

	fmt.Println("The First line is now free")
}

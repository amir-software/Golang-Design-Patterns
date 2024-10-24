package main

import "fmt"

type librarySection struct {
	next Handler
}

func (libhandler *librarySection) setNext(next Handler) {
	libhandler.next = next
}

func (libhandler *librarySection) execute(s *Student) {
	if s.isPassedPaymentSection {
		s.isPassedLibrarySection = true
	} else {
		fmt.Println("The Student Has'n finished his payment")
		return
	}

	fmt.Println("The Student have done it's library stuff")
	libhandler.next.execute(s)
}

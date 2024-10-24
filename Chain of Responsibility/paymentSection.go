package main

import "fmt"

// These are a set (Chain) of handler (Responsibility) for a graduating a student

type paymentSection struct {
	next Handler
}

func (paymenthandler *paymentSection) setNext(next Handler) {
	paymenthandler.next = next
}

func (paymenthandler *paymentSection) execute(s *Student) {
	s.isPassedPaymentSection = true
	fmt.Println("The Student have done it's payment")
	paymenthandler.next.execute(s)
}

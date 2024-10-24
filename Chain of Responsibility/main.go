package main

type Handler interface {
	execute(*Student)
	setNext(Handler)
}

// Handlers in order are :: Payment Library Educational

type Student struct {
	name          string
	lastName      string
	studentNumber int

	isPassedPaymentSection     bool
	isPassedLibrarySection     bool
	isPassedEducationalSection bool
}

func main() {

	student := Student{
		name:          "Max",
		lastName:      "Jackson",
		studentNumber: 12584,
	}

	paymentSectionHandler := paymentSection{}
	librarySectionHandler := librarySection{}
	educationalSectionHandler := educationalSection{}

	paymentSectionHandler.setNext(&librarySectionHandler)
	librarySectionHandler.setNext(&educationalSectionHandler)

	paymentSectionHandler.execute(&student)

}

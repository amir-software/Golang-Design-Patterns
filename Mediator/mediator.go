package main

type Mediator struct {
	isAllow bool

	ListOfUsingStreet []IStreet // Can use this for having the list of the full street lines (Not Implemented)
}

package main

func main() {

	trafficLight := Mediator{isAllow: true}

	streetA := FirstLine{med: &trafficLight}
	streetB := SecondLine{med: &trafficLight}

	if streetA.CanGo() { // Street A is full now
		streetA.Full()
	}

	if streetB.CanGo() { // Street B Can not go because other street is full
		streetB.Full()
	}

	streetA.Free() // Street A is free now

	if streetB.CanGo() { // Street B Can go
		streetB.Full()
	}
	streetB.Free()

}

package main

// Operation InterFaces (it has the contorl of the diffirent behaviors)
type personInterFace interface {
	GetName()
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

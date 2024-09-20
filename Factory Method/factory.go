package main

// Factory Method
func getPerson(kind string) personInterFace {
	if kind == "MALE" {
		return createMale()
	}
	if kind == "FEMALE" {
		return createFeMale()
	}
	return nil
}

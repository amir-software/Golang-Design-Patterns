package main

type Ilaptop interface {
	getInfo()
}

func getClone(brand string) Ilaptop {
	switch brand {

	case "Apple":
		return macbook

	case "Assus":
		return vivobook

	case "Lenovo":
		return ideapad
	}

	return nil
}

func main() {
	laptop := getClone("Assus")
	laptop.getInfo()
}

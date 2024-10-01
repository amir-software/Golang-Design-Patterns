package main

import "fmt"

type Laptop struct {
	brand string
	cpu   string
	ram   string
}

func (lab *Laptop) getInfo() {
	fmt.Printf("The laptop brand is %s and CPU is %s and ram is %s \n", lab.brand, lab.cpu, lab.ram)
}

var macbook *Laptop = &Laptop{
	brand: "Apple",
	cpu:   "8Gen",
	ram:   "3GB",
}

var vivobook *Laptop = &Laptop{
	brand: "Assus",
	cpu:   "12Gen",
	ram:   "12GB",
}

var ideapad *Laptop = &Laptop{
	brand: "Lenovo",
	cpu:   "10Gen",
	ram:   "8GB",
}

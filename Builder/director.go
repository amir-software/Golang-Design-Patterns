package main

// #### Director handle the order and the actions that must be take care of the builing

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director { // Initial function for director
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) { // Set the type of product
	d.builder = b
}

func (d *Director) buildHouse() House { // Create and fetch the final product

	// you can also add condition here and skip or add some steps for diifrenet type of house
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

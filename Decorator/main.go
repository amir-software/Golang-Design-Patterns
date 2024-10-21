package main

import "fmt"



func main (){
	houseA := Apartment{}

	houseAWithStone := stoneDecorator{house: &houseA}

	houseArea := houseAWithStone.getArea()

	fmt.Println(houseArea)

}
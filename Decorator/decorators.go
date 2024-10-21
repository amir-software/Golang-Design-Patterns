package main

type woodDecorator struct{ // Actually decorate the House obj with Wood
	house IHouse
}

func (w *woodDecorator) getArea() int{
	housePrice := w.house.getArea()
	return housePrice + 25
}


type stoneDecorator struct{ // Actually decorate the House obj with Wood
	house IHouse
}

func (w *stoneDecorator) getArea() int{
	housePrice := w.house.getArea()
	return housePrice + 45
}
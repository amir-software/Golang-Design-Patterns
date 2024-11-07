package main

import "fmt"

// This is actually the logic of the programm (Business Rule - Logic)
type LightReceiver struct{}

func (l *LightReceiver) TurnOn() {
	fmt.Println("Light is ON")
}
func (l *LightReceiver) TurnOff() {
	fmt.Println("Light is OFF")
}

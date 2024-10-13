package main

import "fmt"


type printerA struct{}

func (p *printerA) Print(){
	fmt.Println("Printing by Printer type A")
}



type printerB struct{}

func (p *printerB) Print(){
	fmt.Println("Printing by Printer type B")
}



type IPrinter interface{
	Print()
}
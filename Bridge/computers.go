package main

import "fmt"

type MacSystem struct {
	printer IPrinter
}

func (mac *MacSystem) usePrint(){
	fmt.Println("Print would be done by mac system")
	mac.printer.Print()
}


func (mac *MacSystem) SetPrinter(printer IPrinter){
	mac.printer = printer
}



type WindowsSystem struct {
	printer IPrinter
}


func (windows *WindowsSystem) usePrint(){
	fmt.Println("Print would be done by windows system")
	windows.printer.Print()
}

func (windows *WindowsSystem) SetPrinter(printer IPrinter){
	windows.printer = printer
}

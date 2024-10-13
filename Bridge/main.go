package main


func main (){
	
	printerTypeA := printerA{}
	printerTypeB := printerB{}

	macComputer := MacSystem{}

	macComputer.SetPrinter(&printerTypeB)
	macComputer.usePrint()

	macComputer.SetPrinter(&printerTypeA)
	macComputer.usePrint()


	windowsComputer := WindowsSystem{}

	windowsComputer.SetPrinter(&printerTypeB)
	windowsComputer.usePrint()

	windowsComputer.SetPrinter(&printerTypeA)
	windowsComputer.usePrint()
}
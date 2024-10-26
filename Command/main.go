package main


func main (){

	// First step : specify the receiver (Business Logic)
	light := &LightReceiver{}

	// Specify The commands as stand alone objects
	onCommand := &ConcTrunOnCommand{light}
	offCommand := &ConcTrunOffCommand{light}

	// Specify the Invoker and set it a command and then execute the invoker
	remote := &Remote{}
	remote.SetCommandToInvoker(onCommand)
	remote.PressButton()
	remote.SetCommandToInvoker(offCommand)
	remote.PressButton()

}
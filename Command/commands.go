package main

// The Command InterFace
type Command interface{
	execute()
}


// Concrete 'Command' implementation

type ConcTrunOnCommand struct{
	receiver *LightReceiver
}

func (trunon *ConcTrunOnCommand) execute(){
	trunon.receiver.TurnOn()
}


type ConcTrunOffCommand struct{
	receiver *LightReceiver
}

func (trunon *ConcTrunOffCommand) execute(){
	trunon.receiver.TurnOff()
}
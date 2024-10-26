package main



type Remote struct{
	command Command
}


func (r *Remote) SetCommandToInvoker (command Command){
	r.command = command
}

func (r *Remote) PressButton(){
	r.command.execute()
}
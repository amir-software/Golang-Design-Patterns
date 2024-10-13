package main

import "fmt"

type Adapt110to220 struct{
	tv110 *Tv110
}

func (adapter *Adapt110to220) plugInTvWith220Volt(){
	fmt.Println("Converting the 110 to 220 voltage adapter")
	adapter.tv110.plugInTvWith110Volt()	
}

func (adapter *Adapt110to220) getTvType(){
	adapter.tv110.getTvType()
}
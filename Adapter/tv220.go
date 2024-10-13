package main

import "fmt"


type Tv220 struct{
	name string
}

func (tv *Tv220) plugInTvWith220Volt (){
	fmt.Println("the tv had been pluged in to '220' Voltage")
}

func (tv *Tv220) getTvType(){
	fmt.Println(tv.name)
}
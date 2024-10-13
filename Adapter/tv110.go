package main

import "fmt"

type Tv110 struct{
	name string
}

func (tv * Tv110) plugInTvWith110Volt (){
	fmt.Println("the tv had been pluged in to '110' Voltage")
}

func (tv *Tv110) getTvType(){
	fmt.Println(tv.name)
}
package main

import "fmt"

type Singleton struct {
	counter    string
	connection string
}

var singleton *Singleton

func getSingleton() *Singleton {
	if singleton == nil {
		singleton = &Singleton{counter: "0", connection: "0"}
		fmt.Print("Singleton Was Created \n")
		return singleton
	} else {
		fmt.Print("Singleton Was Already Created \n")
		return singleton
	}

}

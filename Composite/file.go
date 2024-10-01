package main

import "fmt"

type File struct {
	name string
}

func (file File) getComponentName() {
	fmt.Printf("the file name is %s \n", file.name)
}

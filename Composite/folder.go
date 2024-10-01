package main

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (folder *Folder) addComponent(c Component) {

	folder.components = append(folder.components, c)
}

func (folder *Folder) getComponentName() {
	fmt.Printf("----- the folder name is %s ----- \n", folder.name)

	for _, comp := range folder.components {
		// fmt.Println("the component name of folder is %s", comp)
		comp.getComponentName()
	}
}

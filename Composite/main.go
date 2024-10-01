package main

func main() {
	file1 := File{name: "File One"}
	file2 := File{name: "File Two"}

	folder1 := Folder{name: "Folder One"}

	folder1.addComponent(file1)
	folder1.addComponent(file2)

	folder1.getComponentName()

}

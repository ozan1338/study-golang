package main

import "log"

type myStruct struct {
	FirstName string
}

//this is called receiver point to the struct so we can using the struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

//Like this we can perform bussines logic

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Marry",
	}

	// log.Println("myVar is set to", myVar.FirstName)
	// log.Println("myVar2 is set to", myVar2.FirstName)

	//now we can call func printFirstName like this
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
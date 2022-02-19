package main

import (
	"log"

	"github.com/ozan1338/myNiceProgram/helpers"
)

func main() {
	log.Println("hello")

	var myVar helpers.SomeType
	myVar.TypeName = "some-name"

	log.Println(myVar.TypeName)
}
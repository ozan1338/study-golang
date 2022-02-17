package main

import "log"

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int //default = 0

	whatToSay, _ = saySomething("Hello") //_ its mean ignore the second param
	log.Println(whatToSay)

	saySomethingElse, _ = saySomething("GoodBye")
	log.Println(saySomethingElse)

	log.Println(saySomething(("Finally")))

	i = 7
	log.Println(i)
}

func saySomething(s string) (string, string) {
	return s, "World"
}
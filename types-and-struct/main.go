package main

import (
	"log"
)

var s = "seven" //Global Variable

func main() {
	var s2 = "six" //local variable inside curly bracket
	//s4 := "eight" //we can declared like this to

	log.Println("s1 is, ", s)
	log.Println("s2 is, ", s2)

	saySomething("xxx")

}

func saySomething(s3 string) (string, string) {
	log.Println("s from saySomething func is,", s)// so actually even we print xxx but the truth is 
	// we override the name var s in the global variable so be careful with name in go
	return s3, "World"
}

//structure 
// type User struct {
// 	FirstName string
// 	LastName string
// 	PhoneNumber string
// 	Age int
// 	BirthDate time.Time
// }
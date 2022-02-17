package main

import (
	"log"
	"time"
)

//structure
//So look like a Object in javascript
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Trevore",
		LastName: "Sawler",
		PhoneNumber: "1-555-555-555",
	}

	log.Println(user.FirstName)
	log.Println(user.BirthDate) //even its not assign in go they have a default already 
	//like int the default is 0
}

//if the letter is start with small letter ist not going to call in another package
//Example func goLang () or var special int
//but if its  start with Capital it can decalred in other package
//Example func GoLang () or var Special int


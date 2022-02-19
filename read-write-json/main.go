package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	HairColor string `json:"hair-color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first-name": "Clark",
			"last-name" : "Kent",
			"hair-color" : "black",
			"has_dog": true
		},
		{
			"first-name": "Bruce",
			"last-name" : "Wayne",
			"hair-color" : "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil{
		log.Println("Error Unmarshaled the json, ", err)
	}

	log.Printf("unmarshalled %v,", unmarshalled)

	//write json from a struct
	var mySlice []Person

	m1 := Person{
		FirstName: "Wally",
		LastName: "West",
		HairColor: "blonde",
		HasDog: true,
	}

	mySlice = append(mySlice, m1)
	
	m2 := Person{
		FirstName: "Diana",
		LastName: "Prince",
		HairColor: "black",
		HasDog: true,
	}

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println(err)
	}

	fmt.Print(string(newJson))
}
package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	//mySlice := []string{"dog","cat","horse","fish","banana"}

	//if we want to loop the array and the first is index and the second is value inside slice
	// for _,x := range mySlice{
	// 	log.Println(x)
	// }

	// myMap := make(map[string]string)
	// myMap["dog"] = "dog"
	// myMap["fish"] = "fish"
	// myMap["horse"] = "horse"

	// for i, x := range myMap {

	// 	log.Println(i, x)
	// }

	var mySlice []User

	user1 := User{
		FirstName: "Trevor",
	}

	user2 := User{
		FirstName: "Sam",
	}

	mySlice = append(mySlice, user1)
	mySlice = append(mySlice, user2)

	for i, x := range mySlice {
		log.Println(i, x.FirstName)
	}

}
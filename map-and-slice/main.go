package main

type User struct {
	FirstName string
	LastName string
}

func main() {
	// myMap := make(map[string]string)
	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"

	//Override map
	//may["dog"] = "Foo"
	// now on key dog its not a samson but foo

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	//myMap1 := make(map[string]int)

	// myMap1["first"] = 1
	// myMap1["second"] = 2

	// log.Println(myMap1["first"])
	// log.Println(myMap1["second"])

	// myMap2 := make(map[string]User)

	// me := User{
	// 	FirstName: "Trevor",
	// 	LastName: "Sawler",
	// }

	// myMap2["me"] = me

	// log.Println(myMap2["me"].FirstName)

	// myMap3 := make(map[int]User)

	// ozan := User{
	// 	FirstName: "Akhmad",
	// 	LastName: "Fauzan",
	// }

	// fadhil := User{
	// 	FirstName: "Fadhil",
	// 	LastName: "Asrofi",
	// }

	// myMap3[0] = ozan
	// myMap3[1] = fadhil

	// for i := 0; i < 2 ; i++ {
	// 	log.Println(myMap3[i])
	// }

	//if you dont know what data type should be inside map
	//just using interface but not recommended
	//myMap := make(map[string]interface{})

	// ========== SLICE ==========
	//in go we almost never use array instead we use slice
	
	// var mySlice []string

	// mySlice = append(mySlice, "Trevor")
	// mySlice = append(mySlice, "john")

	// var mySlice []int

	// mySlice = append(mySlice, 6)
	// mySlice = append(mySlice, 4)

	// sort.Ints(mySlice)

	//We can declare like this too
	// numbers := []int{1, 3, 4, 6, 7}

	// log.Println(numbers)
	// log.Println(numbers[0:2])
	// log.Println(numbers[2:3])

	// ozan := User{
	// 	FirstName: "akhmad",
	// 	LastName: "fauzan",
	// }
	
	// fadhil := User{
	// 	FirstName: "fadhil",
	// 	LastName: "asrofi",
	// }
	
	// var names = []User{ozan,fadhil}

	// log.Println(names[0].FirstName)
}
package main

import "log"

func main() {
	// isTrue := true

	// if isTrue == false {
	// 	log.Println(isTrue)
	// }else {
	// 	//same as javascript put ! is making kebalikannya dari
	// 	//nilai tersebut
	// 	log.Println(!isTrue)
	// }
	
	//myNum,isTrue := 100,false

	// if myNum > 99 && !isTrue {
	// 	log.Println("myNum is greater than 99 and var isTrue is true")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("3")
	// } else if myNum == 101 || isTrue {
	// 	log.Println("2")

	// } else if myNum > 1000 && !isTrue {
	// 	log.Println("1")

	// }

	myVar := "cat"
	myVar = "horse"

	switch myVar {
	case "cat" :
		log.Println(myVar)
		log.Println("heyoo")

	case "dog" :
		log.Println(myVar)

	case "fish" :
		log.Println(myVar)
	
	default: 
	log.Println("something else")
	}
}
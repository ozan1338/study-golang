package main

import "log"

func main () {
	var myString string = "Grenn"

	log.Println("my String is, ", myString)
	changeMyString(&myString)
	log.Println("after func called my String is, ", myString)
}

func changeMyString (s *string) {
	log.Println("s is, ", s)
	newValue := "Red"
	*s = newValue
}
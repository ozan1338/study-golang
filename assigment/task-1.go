package main

import "fmt"

func main () {
	var firstName string
	var lastName string

	fmt.Printf("Input Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Input Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Your Name: ", firstName+lastName)
}
package main

import (
	"fmt"
	"net/http"
)

//Create Custom Error
//errors.New("whatever")

const PORT = ":8080"

//main is the main application func
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	_,_ = fmt.Println("Starting App on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
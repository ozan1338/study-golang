package main

import (
	"fmt"
)

func main(){
	num := 10
	if num < 10 {
		fmt.Printf("Less Than 10")
	}else if num > 10{
		fmt.Printf("More Than 10")
	}else{
		fmt.Printf("Exactly 10")
	}

	//refactor if else
	if num <= 10 {
		fmt.Printf("Less or Equal than 10")
	}else if num > 10 {
		fmt.Printf("More than 10")
	}
}
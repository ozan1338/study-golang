package main

import (
	"fmt"
	"math"
)

func main(){
	var1 := math.Pi
	var2 := "Hey Gopher"
	var3 := true
	var4 := 49
	
	fmt.Printf("Value: %.2f",var1)
	fmt.Printf("\tData Type: %T\n",var1)
	fmt.Printf("Value: %v",var2)
	fmt.Printf("\tData Type: %T\n",var2)
	fmt.Printf("Value: %v",var3)
	fmt.Printf("\tData Type: %T\n",var3)
	fmt.Printf("Value: %v",var4)
	fmt.Printf("\tData Type: %T\n",var4)
}
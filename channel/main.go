package main

import (
	"log"

	"github.com/ozan1338/study-go/helpers"
)

func main() {
	intChan := make(chan int)
	defer close(intChan)
	//defer meant is close as soon the func is done

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}

const numPool = 100

func CalculateValue (intChan chan int) {
	randomNumber := helpers.RandNumber(numPool)
	intChan <- randomNumber
}


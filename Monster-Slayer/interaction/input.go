package interaction

import (
	"fmt"
)

// var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialRound bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1"  {
			return "ATTACK"
		} else if playerChoice == "2"  {
			return "HEAL"
		} else if playerChoice == "3" {
			if isSpecialRound {
				return "SPECIAL ATTACK"
			} else {
				return "IS NOT TURN FOR SPECIAL ATTACK"
			}
		}

		fmt.Println("Fetching The user input failed. Please Try Again!!")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your Choice : ")

	var userInput string

	fmt.Scanln(&userInput)

	// if err != nil{
	// 	return "", err
	// }


	return userInput, nil
}
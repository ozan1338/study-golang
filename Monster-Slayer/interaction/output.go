package interaction

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action string
	PlayerAttackDamage int
	PlayerHealthValue int
	MonsterAttackDamage int
	PlayerHealth int
	MonterHealth int
}



func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting the Game...")
	fmt.Println("Good Luck!!")
}

func ShowAvailableActions(isSpecialRound bool) {
	fmt.Println("Please Choose Your Action")
	fmt.Println("-------------------------")
	fmt.Println("1). Attack Monster")
	fmt.Println("2). Heal")

	if isSpecialRound {
		fmt.Println("3). Special Attack")
	}
}

func PrintRoundStatistic(roundData *RoundData) {
	switch roundData.Action {
	case "ATTACK" :
		fmt.Printf("Player Attacked monster for %v damage\n", roundData.PlayerAttackDamage)
	case "SPECIAL ATTACK":
		fmt.Printf("Player Attacked with Special Attack monster for %v damage\n", roundData.PlayerAttackDamage)
	case "HEAL":
		fmt.Printf("Player Healed for %v.\n", roundData.PlayerHealthValue)
	}

	fmt.Printf("Monster attacked player for %v damage\n", roundData.MonsterAttackDamage)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("-------------------------")
	fmt.Printf("%v WON!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file,err := os.Create("gameLog.txt")

	if err != nil{
		fmt.Println("Saving a log file failed. Exiting...")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"RoundKey": fmt.Sprint(index+1),
			"Action": value.Action,
			"Player Attack Damage": fmt.Sprint(value.PlayerAttackDamage),
			"Player Health Value": fmt.Sprint(value.PlayerHealthValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDamage),
			"Player Health": fmt.Sprint(value.PlayerHealth),
			"Monster Health": fmt.Sprint(value.MonterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log File Failed")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote Data To Log!")
}
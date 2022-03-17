package main

import (
	"github.com/ozan1338/monsterslayer/actions"
	"github.com/ozan1338/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++

	isSpecialRound := currentRound % 3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealthValue int
	var monsterAttackDmg int

	switch userChoice {
	case "ATTACK":
		playerAttackDmg = actions.AttackMonster(false)
	case "HEAL":
		playerHealthValue = actions.HealPlayer()
	case "SPECIAL ATTACK":
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmount()

	roundData := interaction.RoundData{
		Action: userChoice,
		PlayerHealth: playerHealth,
		MonterHealth: monsterHealth,
		PlayerAttackDamage: playerAttackDmg,
		MonsterAttackDamage: monsterAttackDmg,
		PlayerHealthValue: playerHealthValue,
	}

	interaction.PrintRoundStatistic(&roundData)
	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	}else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var currentMosterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMosterHealth -= dmgValue

	return dmgValue
}

func HealPlayer() int {
	minHealthValue := PLAYER_HEALTH_MIN_VALUE
	maxHealthValue := PLAYER_HEALTH_MAX_VALUE

	healthValue := generateRandBetween(minHealthValue, maxHealthValue)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healthValue {
		currentPlayerHealth += healthValue
		return healthValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTCK_MIN_DMG
	maxAttackValue := MONSTER_ATTCK_MAX_DMG

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dmgValue

	return dmgValue
}

func GetHealthAmount() (int,int) {
	return currentPlayerHealth,currentMosterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max - min) + min
}
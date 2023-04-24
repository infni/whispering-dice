package rolling

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RollDamageMessage(diceCount int, dieCap int) string {

	// these setting are checked prior to callign this function, these checks are just protection to prevent accidental infinte loops.
	if dieCap < 1 {
		dieCap = 1
	}
	if diceCount < 1 {
		diceCount = 1
	}

	singleDieRolls := make([]int, 0)
	completedRollCount := 0
	total := 0
	for {
		roll := random.Intn(6) + 1
		singleDieRolls = append(singleDieRolls, roll)
		if roll <= dieCap {
			completedRollCount++
			total += roll
			if completedRollCount >= diceCount {
				break
			}
		}
	}

	// format
	formattedRolls := make([]string, len(singleDieRolls))
	for i, val := range singleDieRolls {
		if val > dieCap {
			formattedRolls[i] = fmt.Sprintf("~~_%d_~~", val)
		} else {
			formattedRolls[i] = fmt.Sprintf("%d", val)
		}
	}

	return fmt.Sprintf("Damage: **%d** Cap: %d [%s]", total, dieCap, strings.Join(formattedRolls, ","))
}

func RollChallengeMessage(dieCount int, skill int) string {

	rolls := make([]int, dieCount)
	for i := 0; i < len(rolls); i++ {
		rolls[i] = random.Intn(6) + 1
	}

	maxSingleDie := 1
	for i := 0; i < len(rolls); i++ {
		if maxSingleDie < rolls[i] {
			maxSingleDie = rolls[i]
		}
	}
	totals := make([]int, 6)
	for i := 0; i < len(rolls); i++ {
		totals[rolls[i]-1] += rolls[i]
	}
	maxRoll := maxSingleDie
	maxRollSide := maxSingleDie
	for i := 0; i < len(totals); i++ {
		if totals[i] > maxRoll {
			maxRoll = totals[i]
			maxRollSide = i + 1
		}
	}

	formattedDice := make([]string, len(rolls))
	for i := 0; i < len(rolls); i++ {
		if rolls[i] == maxRollSide {
			formattedDice[i] = fmt.Sprintf("**%d**", rolls[i])
		} else {
			formattedDice[i] = fmt.Sprintf("%d", rolls[i])
		}
	}

	return fmt.Sprintf("[%s] Roll: **%d**", strings.Join(formattedDice, ","), maxRoll)
}

package rolling

func RollChallengeAllRolls(diceCount int) (rolls []int, maxRoll int, selectedSide int) {

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
	maxRoll = maxSingleDie
	selectedSide = maxSingleDie
	for i := 0; i < len(totals); i++ {
		if totals[i] > maxRoll {
			maxRoll = totals[i]
			selectedSide = i + 1
		}
	}

	return
}

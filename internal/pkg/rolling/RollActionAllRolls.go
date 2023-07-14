package rolling

func RollActionAllRolls(diceCount int, db int, as int, reroll map[int]struct{}) (rolls []int, total int) {

	rolls = make([]int, 0)
	countedDice := 0
	total = 0
	for countedDice < diceCount {

		roll := random.Intn(10) + 1
		rolls = append(rolls, roll)
		if _, ok := reroll[roll]; ok {
			continue
		}
		countedDice++

		if roll >= db {
			total += 2
		} else if roll >= 7 {
			total += 1
		}

	}

	total += as

	return rolls, total
}

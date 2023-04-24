package rolling

func RollDamageAllRolls(diceCount int, dieCap int) (rolls []int, total int) {

	// these setting are checked prior to callign this function, these checks are just protection to prevent accidental infinte loops.
	if dieCap < 1 {
		dieCap = 1
	}
	if diceCount < 1 {
		diceCount = 1
	}

	rolls = make([]int, 0)
	total = 0
	completedRollCount := 0
	for {
		roll := random.Intn(6) + 1
		rolls = append(rolls, roll)
		if roll <= dieCap {
			completedRollCount++
			total += roll
			if completedRollCount >= diceCount {
				break
			}
		}
	}

	return rolls, total
}

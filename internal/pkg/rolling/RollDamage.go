package rolling

func RollDamage(diceCount int, dieCap int) (rolls []int, total int) {

	rolls = make([]int, diceCount)
	total = 0
	for i := 0; i < len(rolls); i++ {
		rolls[i] = random.Intn(6) + 1
		if rolls[i] <= dieCap {
			total += rolls[i]
		}
	}

	return rolls, total
}

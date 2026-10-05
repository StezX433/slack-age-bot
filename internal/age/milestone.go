package age

func NextMilestone(currentAge int) int {
	if currentAge < 25 {
		return 25
	}

	return ((currentAge / 5) + 1) * 5
}

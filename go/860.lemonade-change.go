func lemonadeChange(bills []int) bool {
	var five, ten, twenty int

	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			five--
			ten++
		} else {
			twenty++
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if !(twenty >= 0 && ten >= 0 && five >= 0) {
			return false
		}
	}

	return true
}

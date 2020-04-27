func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if tank := gas[i] - cost[i]; tank >= 0 && canComplete(gas, cost, tank, (i+1)%len(gas), i) {
			return i
		}
	}
	return -1
}

func canComplete(gas, cost []int, tank, i, end int) bool {
	if i == end {
		return true
	}
	if tank = tank + gas[i] - cost[i]; tank >= 0 {
		return canComplete(gas, cost, tank, (i+1)%len(gas), end)
	} else {
		return false
	}
}
/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// smarter solution is omitted
	// https://leetcode.com/problems/gas-station/discuss/42568/Share-some-of-my-ideas
	// 1. if car starts at A and can not reach B. any station between A and B
	// 2. if the total number of gas is bigger than the total number of cost. there must be a solution. 
	return naive(gas, cost)
}

// O(n^2)
func naive(gas, cost []int) int {
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
// @lc code=end


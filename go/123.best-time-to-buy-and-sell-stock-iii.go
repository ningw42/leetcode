/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
func maxProfit(prices []int) int {
	left := make([]int, len(prices))
	right := make([]int, len(prices))
	var sum, max, min int
	for i := 1; i < len(prices); i++ {
		if newSum := sum + prices[i] - prices[i-1]; newSum < 0 {
			sum = 0
		} else {
			sum = newSum
		}
		if sum > max {
			max = sum
		}
		left[i] = max
	}
	sum = 0
	for i := len(prices) - 2; i >= 0; i-- {
		if newSum := sum + prices[i] - prices[i+1]; newSum > 0 {
			sum = 0
		} else {
			sum = newSum
		}
		if sum < min {
			min = sum
		}
		right[i] = -min
	}

	for i := 0; i < len(prices); i++ {
		if sum := left[i] + right[i]; sum > max {
			max = sum
		}
	}

	return max
}


// AC with O(n^2)
func naive(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if profit := maxProfitOnce(prices[:i]) + maxProfitOnce(prices[i:]); profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
func maxProfitOnce(prices []int) int {
	var sum, max int
	for i := 1; i < len(prices); i++ {
		if newSum := sum + (prices[i] - prices[i-1]); newSum < 0 {
			sum = 0
		} else {
			sum = newSum
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
// @lc code=end


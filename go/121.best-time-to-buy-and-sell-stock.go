/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
func maxProfit(prices []int) int {
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


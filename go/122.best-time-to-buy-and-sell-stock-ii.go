/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if todayProfit := prices[i] - prices[i-1]; todayProfit > 0 {
			profit += todayProfit
		}
	}
	return profit
}
// @lc code=end


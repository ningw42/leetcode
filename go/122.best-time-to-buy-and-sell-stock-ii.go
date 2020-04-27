func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if todayProfit := prices[i] - prices[i-1]; todayProfit > 0 {
			profit += todayProfit
		}
	}
	return profit
}
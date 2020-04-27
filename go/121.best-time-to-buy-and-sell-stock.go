func maxProfit(prices []int) int {
    profits := make([]int, len(prices))
    var maxProfit int
    for i := 1; i < len(prices); i++ {
        todayProfit := profits[i-1] + prices[i] - prices[i-1]
        profits[i] = max(0, todayProfit)
        if profits[i] > maxProfit {
            maxProfit = profits[i]
        }
    }
    return maxProfit
}

func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}
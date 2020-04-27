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

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	sums := make(map[int]bool)
	sums[n] = true
	sum := n
	for {
		sum = squareSum(sum)
		if sum == 1 {
			return true
		}
		if _, exists := sums[sum]; exists {
			return false
		}
		sums[sum] = true
	}
}

func squareSum(n int) int {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
// @lc code=end


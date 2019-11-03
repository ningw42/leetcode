/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	factors := []int{2, 3, 5}
	results := make([]int, n)
	results[0] = 1

	for i := 1; i < n; i++ {
		next := 0
		for j := 0; j < i; j++ {
			for _, factor := range factors {
				if product := results[j] * factor; product > results[i-1] {
					if next == 0 {
						next = product
					} else {
						if product < next {
							next = product
						}
					}
				}
			}
		}
		results[i] = next
	}

	return results[n-1]
}

// @lc code=end


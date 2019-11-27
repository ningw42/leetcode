/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

// TODO better time complexity
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	lis := make([][][]int, len(nums))
	lis[0] = [][]int{[]int{nums[0]}}
	maxLength := 0
	for i := 1; i < len(nums); i++ {
		var results [][]int
		for j := 0; j < i; j++ {
			prev := lis[i-1][j]
			var result []int
			if maxElem := prev[len(prev)-1]; maxElem < nums[i] {
				result = append(prev, nums[i])
			} else {
				if index := len(prev) - 2; index >= 0 && prev[index] < nums[i] {
					result = append(prev[0 : index + 1], nums[i])
				} else {
					result = prev
				}
			}
			if length := len(result); length > maxLength {
				maxLength = length
			}
			results = append(results, result)
		}
		results = append(results, []int{nums[i]})
		lis[i] = results
	}

	return maxLength
}
// @lc code=end


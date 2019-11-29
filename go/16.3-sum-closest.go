/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	var result int
	var abs float64
	abs = float64(math.MaxInt32)
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		// the first element of result fixed
		// skip all elements with the same value
		if i == 0 || nums[i-1] != nums[i] {
			t := target - nums[i]
			j := i + 1
			k := len(nums) - 1
			for j < k {
				sum := nums[j] + nums[k]
				if newAbs := math.Abs(float64(sum - t)); newAbs < abs {
					abs = newAbs
					result = sum + nums[i]
				}
				if sum > t {
					// skip elements with same value
					for j < k && nums[k-1] == nums[k] {k--}
					k--
				} else if sum < t {
					// skip elements with same value
					for j < k && nums[j+1] == nums[j] {j++}
					j++
				} else {
					return sum + nums[i]
				}
			}
		}
	}

	return result 
}
// @lc code=end


/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	} 
	if len(nums) == 1 {
		return nums[0] == target
	}

	length := len(nums)
	var left, right, mid int
	for left = 1; left < length; left++ {
		if nums[left] < nums[left-1] {
			break
		}
	}
	left = left % length
	right = (left - 1 + length) % length

	fmt.Println(left, right)
	for left != right {
		if left < right {
			mid = (left + right) / 2
		} else {
			mid = (left + right + length) / 2 % length
		}
		if val := nums[mid]; val > target {
			right = mid
			// if left < right {
			// 	right = mid - 1
			// } else {
			// 	right = (mid - 1 + length) % length
			// }
		} else if val < target {
			if left < right {
				left = mid + 1
			} else {
				left = (mid + 1 + length) % length
			}
		} else {
			return true
		}
	}

	return nums[left] == target
}
// @lc code=end


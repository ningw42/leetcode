/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
func search(nums []int, target int) int {
	// there must be at least ons side of the divided array is sorted
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		// target found
		if nums[mid] == target {
			return mid
		}

		// left side is sorted
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				// target fits in the sorted left side
				right = mid - 1
			} else {
				// target fits in the right side
				left = mid + 1
			}
		}

		// right side is sorted
		if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				// target fits in the sorted right side
				left = mid + 1
			} else {
				// target fits in the left side
				right = mid - 1
			}
		}
	}

	return -1
}


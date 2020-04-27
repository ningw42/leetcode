/*
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (39.17%)
 * Total Accepted:    1.2M
 * Total Submissions: 3.1M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 *
 *
 */
func twoSum(nums []int, target int) []int {
	return twoSumMap(nums, target)
}

func twoSumMap(nums []int, target int) []int {
	m := make(map[int]int)

	for pos, value := range nums {
		subTarget := target - value
		if index, exist := m[subTarget]; exist {
			return []int{pos, index}
		} else {
			m[value] = pos
		}
	}

	return []int{}
}

func twoSumNaive(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

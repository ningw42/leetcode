import (
	"sort"
	"strconv"
)

/*
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (22.22%)
 * Total Accepted:    415.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
func threeSum(nums []int) [][]int {
	return threeSumSorted(nums)
}

// may exceed the memory limit
func threeSumSorted(nums []int) [][]int {
	var result [][]int

	// sort original slice asc
	sort.Ints(nums)

	for i := 0; i <= len(nums)-2; i++ {
		a := nums[i]
		// the first element is already grater than zero, break
		if a > 0 {
			break
		}
		// first element in current iteration is identical to the last iteration, skip current iteration
		if i > 0 && a == nums[i-1] {
			continue
		}

		// searching from both side
		j := i + 1
		k := len(nums) - 1
		for j < k {
			b := nums[j]
			c := nums[k]
			sum := a + b + c
			if sum == 0 {
				// found the triplet
				result = append(result, []int{a, b, c})
				// move the left cursor to the next value with different value
				for j = j + 1; j < k && b == nums[j]; j = j + 1 {
				}
				// move the right cursor to the next value with different value
				for k = k - 1; j < k && c == nums[k]; k = k - 1 {
				}
			} else if sum < 0 {
				for j = j + 1; j < k && b == nums[j]; j = j + 1 {
				}
			} else {
				for k = k - 1; j < k && c == nums[k]; k = k - 1 {
				}
			}
		}
	}

	return result
}

// won't AC, O(n^3)
func threeSumNaiveTraversal(nums []int) [][]int {
	m := make(map[string]bool)
	var result [][]int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					s := []int{nums[i], nums[j], nums[k]}
					uniqueKey := genUniqueKey(s)
					if _, exist := m[uniqueKey]; !exist {
						result = append(result, s)
						m[uniqueKey] = true
					}
				}
			}
		}
	}

	return result
}

func genUniqueKey(a []int) string {
	sort.Ints(a)
	ret := ""
	for _, num := range a {
		ret = ret + strconv.Itoa(num) + ","
	}

	return ret
}

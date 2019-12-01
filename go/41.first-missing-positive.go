/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start

// https://leetcode.com/problems/first-missing-positive/discuss/17080/Python-O(1)-space-O(n)-time-solution-with-explanation

// 1. for any array whose length is l, the first missing positive must be in range [1,...,l+1], 
// so we only have to care about those elements in this range and remove the rest.
// 2. we can use the array index as the hash to restore the frequency of each number within 
//  the range [1,...,l+1] 

//  after removing all the numbers greater than or equal to n, all the numbers remaining are smaller than n. If any number i appears, we add n to nums[i] which makes nums[i]>=n. Therefore, if nums[i]<n, it means i never appears in the array and we should return i.
func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	length := len(nums) // length is also the max possible answer
	for i := 0; i < length; i++ {
		// all elements greater than should be removed
		// if element 'length' is present, the range is restricted to [1, length-1]
		// so nums[i] >= length
		if nums[i] < 0 || nums[i] >= length {
			nums[i] = 0
		}
	}
	for i := 0; i < length; i++ {
		nums[nums[i]%length] += length
	}
	for i := 1; i < length; i++ {
		if nums[i] < length {
			return i
		}
	}

	return length 
}

// O(n) time complexity but O(n) space complexity
func naive(nums []int) int {
	all := make(map[int]bool)
	for _, num := range nums {
		if num > 0 {
			all[num] = true
		}
	}

	i := 1
	for {
		if _, exists := all[i]; !exists {
			break
		}
		i++
	}

	return i
}
// @lc code=end


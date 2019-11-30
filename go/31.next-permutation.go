/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	lnums := sort.IntSlice(nums)
	var i int
    for i = lnums.Len() - 2; i >= 0; i-- {
		if lnums.Less(i, i+1) {
			break
		}
	}

	if i < 0 {
		// sort.Reverse() produces a new slice
		reverseInPlace(lnums)
	} else {
		var index int
		abs := math.MaxInt32
		for j := i + 1; j < lnums.Len(); j++ {
			if lnums.Less(i, j) {
				if newAbs := nums[j] - nums[i]; newAbs < abs {
					index = j
					abs = newAbs
				}
			}
		}
		lnums.Swap(i, index)
		tail := sort.IntSlice(lnums[i+1:])
		sort.Sort(tail)
	}
}

func reverseInPlace(nums sort.IntSlice) {
	for i := 0; i < nums.Len() / 2; i++ {
		nums.Swap(i, nums.Len() - 1 - i)
	}
}
// @lc code=end


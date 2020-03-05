/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start

var result [][]int
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result = [][]int{}
	findUniquePermute([]int{}, nums)

	return result
}

// No.1
func permutationUnique(nums []int, taken map[int]struct{}) [][]int {
    if len(nums) == len(taken) {
        return [][]int{nil}
    }
    var ret [][]int
    for i := 0; i < len(nums); {
        if _, isTaken := taken[i]; !isTaken {
            taken[i] = struct{}{}
            for _, p := range permutationUnique(nums, taken) {
                ret = append(ret, append([]int{nums[i]}, p...))
            }
            delete(taken, i)
            for i = i + 1; i < len(nums) && nums[i] == nums[i-1]; i++ {}
        } else {
            i++
        }
    }
    return ret
}

// No.2
func findUniquePermute(permute []int, nums []int) {
	if len(nums) == 0 {
		result = append(result, makeSliceCopy(permute))
	} else {
		for i := 0; i < len(nums); {
			findUniquePermute(append(permute, nums[i]), removeElementFromSlice(nums, i))
			for i = i + 1; i < len(nums) && nums[i] == nums[i-1]; i++ {}
		}
	}
}

func makeSliceCopy(s []int) []int {
	copy := make([]int, len(s))
	for i, e := range s {
		copy[i] = e
	}
	return copy
}

func removeElementFromSlice(s []int, i int) []int {
	var copy []int
	return append(append(copy, s[0:i]...), s[i+1:]...)
}

// @lc code=end


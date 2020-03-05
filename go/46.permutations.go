/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

func permute(nums []int) [][]int {
    return permutation(nums, make(map[int]struct{}))
}

func permutation(nums []int, taken map[int]struct{}) [][]int {
    if len(nums) == len(taken) {
        return [][]int{nil}
    }
    var ret [][]int
    for i, num := range nums {
        if _, isTaken := taken[i]; !isTaken {
            taken[i] = struct{}{}
            for _, p := range permutation(nums, taken) {
                ret = append(ret, append([]int{num}, p...))
            }
            delete(taken, i)
        }
    }
    return ret
}
// @lc code=end


func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    return permutationUnique(nums, make(map[int]struct{}))
}

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
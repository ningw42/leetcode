/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
    return naive(nums)
}

func naive(nums []int) int {
    length := len(nums)
    start := 0
    count := 0
    for i := 1; i < length; i++ {
        if nums[start] == nums[i] {
            count++
        } else {
            if count != 0 {
                shiftLeft(nums, length, start+1, count)
                length -= count
                i -= count
            }
            start = i
            count = 0
        }
    }
    if count != 0 {
        shiftLeft(nums, length, start+1, count)
        length -= count
    }
    return length
}

func shiftLeft(nums []int, length int, start int, count int) {
    for i := start; i+count < length; i++ {
        nums[i] = nums[i+count]
    }
}


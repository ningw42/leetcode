// time: O(N)
// space: O(N)
// slow/fast pointer should reduce space complexity to O(1)
func circularArrayLoop(nums []int) bool {
    result := make(map[int]bool)
    // defer fmt.Println(result)
    for i := 0; i < len(nums); i++ {
        if _, exists := result[i]; !exists {
            r := try(nums, map[int]int{}, i + nums[i], nums[i] > 0, result)
            if r {
                return true
            }
            result[i] = r
        }
    }
    return false
}

func try(nums []int, history map[int]int, current int, direction bool, result map[int]bool) bool {
    current = current % len(nums)
    if current < 0 {
        current += len(nums)
    }
    if index, traversed := history[current]; traversed {
        return len(history) - index > 1
    } else {
        if (nums[current] > 0) != direction {
            // not the same direction
            // return false but not mark current
            return false
        } else {
            // same direction
            history[current] = len(history)
            ret := try(nums, history, current + nums[current], nums[current] > 0, result)
            result[current] = ret
            return ret
        }
    }
}
